package logger

import (
	"context"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/blendle/zapdriver"
	"go.uber.org/zap"
)

func Middleware(ctx context.Context, logger *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("X-Cloud-Trace-Context")

			if len(header) > 0 {
				traceID, spanID, sampled := deconstructXCloudTraceContext(header)

				log.Printf("trace: %s, spanID: %s", traceID, spanID)

				fields := zapdriver.TraceContext(traceID, spanID, sampled, os.Getenv("GCP_PROJECT_ID"))
				logger = logger.With(fields...)
			}

			next.ServeHTTP(w, r)
		})
	}

}

var reCloudTraceContext = regexp.MustCompile(
	// Matches on "TRACE_ID"
	`([a-f\d]+)?` +
		// Matches on "/SPAN_ID"
		`(?:/([a-f\d]+))?` +
		// Matches on ";0=TRACE_TRUE"
		`(?:;o=(\d))?`)

func deconstructXCloudTraceContext(s string) (traceID, spanID string, traceSampled bool) {
	matches := reCloudTraceContext.FindStringSubmatch(s)

	traceID, spanID, traceSampled = matches[1], matches[2], matches[3] == "1"

	if spanID == "0" {
		spanID = ""
	}

	return
}
