package request

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
)

const (
	HeaderRequestID = "X-Request-ID"
	HeaderTraceID   = "X-Trace-ID"
)

// Middleware 建立 Edge Request Context。
//
// HTTP:
//
//	Request
//	   ↓
//	X-Request-ID / X-Trace-ID
//	   ↓
//	request.Context
//	   ↓
//	context.Context
func Middleware() func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {

				requestID := r.Header.Get(HeaderRequestID)
				if requestID == "" {
					requestID = newID()
				}

				traceID := r.Header.Get(HeaderTraceID)
				if traceID == "" {
					traceID = newID()
				}

				requestContext := Context{
					RequestID: requestID,
					TraceID:   traceID,
				}

				ctx := WithContext(
					r.Context(),
					requestContext,
				)

				// 将最终使用的 ID 返回给 Consumer。
				w.Header().Set(HeaderRequestID, requestID)
				w.Header().Set(HeaderTraceID, traceID)

				next.ServeHTTP(w, r.WithContext(ctx))
			})
	}
}

func newID() string {

	buffer := make([]byte, 16)

	if _, err := rand.Read(buffer); err != nil {
		// crypto/rand 在正常运行环境下不应该失败。
		// v0.1 保持实现简单。
		panic(err)
	}

	return hex.EncodeToString(buffer)
}

// abortInternalServerError 是预留给后续 request middleware
// 更复杂错误处理时使用的最小 HTTP 行为。
func abortInternalServerError(w http.ResponseWriter) {
	http.Error(
		w,
		"internal server error",
		http.StatusInternalServerError,
	)
}
