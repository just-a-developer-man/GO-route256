package middleware

import (
	"fmt"
	"log/slog"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/just-a-developer-man/GO-route256/workshop-1/loms/internal/logging"
)

// WithHTTPRecoverMiddleware recover panics
func WithHTTPRecoverMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				slog.InfoContext(r.Context(),
					"err", err,
					"trace", debug.Stack(),
				)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

// AddLoggingCtxMiddleware adds logging context to request context
func AddLoggingCtxMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		reqCtx := r.Context()
		reqCtx = logging.WithLogRequest(reqCtx, r.Method, r.URL.Path)
		r = r.WithContext(reqCtx)
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

// LogHandlingTimeAndStatusMiddlware logs request handling status code and time
func LogHandlingTimeAndStatusMiddlware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		customRW := &customResponseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}
		next.ServeHTTP(customRW, r)
		handleTime := time.Since(start)
		slog.DebugContext(r.Context(), "request handled", "status", customRW.statusCode, "handle time", fmt.Sprintf("%d ns", handleTime))
	}
	return http.HandlerFunc(fn)
}

type customResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *customResponseWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}
