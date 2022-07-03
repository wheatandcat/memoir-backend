package logger

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"go.uber.org/zap"
)

func Middleware(ctx context.Context, logger *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			traceHeader := r.Header.Get("X-Cloud-Trace-Context")
			traceParts := strings.Split(traceHeader, "/")
			zap.L().Info("log", zap.String("traceHeader", traceHeader))
			if len(traceParts) > 0 {
				traceId := traceParts[0]
				logger.With(
					zap.String("logging.googleapis.com/trace", fmt.Sprintf("projects/%s/traces/%s", os.Getenv("GCP_PROJECT_ID"), traceId)),
				)
			}

			next.ServeHTTP(w, r)
		})
	}

}
