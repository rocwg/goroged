package request

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"

	"github.com/gin-gonic/gin"
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
func Middleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		requestID := c.GetHeader(HeaderRequestID)
		if requestID == "" {
			requestID = newID()
		}

		traceID := c.GetHeader(HeaderTraceID)
		if traceID == "" {
			traceID = newID()
		}

		requestContext := Context{
			RequestID: requestID,
			TraceID:   traceID,
		}

		ctx := WithContext(
			c.Request.Context(),
			requestContext,
		)

		c.Request = c.Request.WithContext(ctx)

		// 将最终使用的 ID 返回给 Consumer。
		c.Header(HeaderRequestID, requestID)
		c.Header(HeaderTraceID, traceID)

		c.Next()
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
func abortInternalServerError(c *gin.Context) {
	c.AbortWithStatusJSON(
		http.StatusInternalServerError,
		gin.H{
			"code":    "INTERNAL_SERVER_ERROR",
			"message": "internal server error",
		},
	)
}
