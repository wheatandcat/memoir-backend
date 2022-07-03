package logger

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"

	"go.uber.org/zap"
)

func Middleware(ctx context.Context, logger *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			traceHeader := r.Header.Get("X-Cloud-Trace-Context")
			traceParts := strings.Split(traceHeader, "/")
			if len(traceParts) > 0 && traceParts[0] != "" {
				trace, spanID := deconstructXCloudTraceContext(traceParts[0])
				logger = logger.With(
					zap.String("logging.googleapis.com/trace", fmt.Sprintf("projects/%s/traces/%s", os.Getenv("GCP_PROJECT_ID"), trace)),
					zap.String("logging.googleapis.com/spanId", spanID),
				)
			}

			zap.ReplaceGlobals(logger)
			next.ServeHTTP(w, r)
		})
	}

}

var reCloudTraceContext = regexp.MustCompile(`([a-f\d]+)/([a-f\d]+)`)

func deconstructXCloudTraceContext(s string) (traceID, spanID string) {
	matches := reCloudTraceContext.FindAllStringSubmatch(s, -1)
	if len(matches) != 1 {
		return
	}

	sub := matches[0]
	if len(sub) != 3 {
		return
	}

	traceID, spanID = sub[1], sub[2]
	if spanID == "0" {
		spanID = ""
	}

	return
}
