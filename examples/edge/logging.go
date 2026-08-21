package main

import (
	"log"
	"net/http"
	"time"

	gedidentity "github.com/rocwg/ged/identity"
	gedrequest "github.com/rocwg/ged/request"
)

type statusWriter struct {
	http.ResponseWriter
	status int
}

func (w *statusWriter) WriteHeader(status int) {
	if w.status != 0 {
		return
	}
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *statusWriter) Write(body []byte) (int, error) {
	if w.status == 0 {
		w.WriteHeader(http.StatusOK)
	}
	return w.ResponseWriter.Write(body)
}

func LoggingMiddleware() func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {

				start := time.Now()

				sw := &statusWriter{
					ResponseWriter: w,
				}

				next.ServeHTTP(sw, r)

				requestContext, _ := gedrequest.FromContext(r.Context())
				identityContext, _ := gedidentity.FromContext(r.Context())

				status := sw.status
				if status == 0 {
					status = http.StatusOK
				}

				log.Printf(
					"request completed method=%s path=%s status=%d duration=%s request_id=%s trace_id=%s user_id=%s tenant_id=%s",
					r.Method,
					r.URL.Path,
					status,
					time.Since(start),
					requestContext.RequestID,
					requestContext.TraceID,
					identityContext.UserID,
					identityContext.TenantID,
				)
			},
		)
	}
}
