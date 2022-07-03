package logger

import (
	"context"
	"net/http"
	"os"
	"regexp"

	"github.com/blendle/zapdriver"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

func Middleware(ctx context.Context, logger *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sc := trace.SpanContextFromContext(r.Context())
			traceID := sc.TraceID().String()
			spanID := sc.SpanID().String()

			zap.L().Info("request", zap.String("trace_id", traceID), zap.String("span_id", spanID))

			if sc.IsValid() {
				fields := zapdriver.TraceContext(sc.TraceID().String(), sc.SpanID().String(), sc.IsSampled(), os.Getenv("GCP_PROJECT_ID"))
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
