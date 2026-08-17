package observability

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/gin-gonic/gin"

	edgerequest "github.com/rocwg/goro-edge/runtime/request"
)

const (
	HeaderRequestID = "X-Request-ID"
	HeaderTraceID   = "X-Trace-ID"
)

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

		requestContext := edgerequest.Context{
			RequestID: requestID,
			TraceID:   traceID,
		}

		ctx := edgerequest.WithContext(
			c.Request.Context(),
			requestContext,
		)

		c.Request = c.Request.WithContext(ctx)

		c.Header(
			HeaderRequestID,
			requestID,
		)

		c.Header(
			HeaderTraceID,
			traceID,
		)

		c.Next()
	}
}

func newID() string {
	buffer := make([]byte, 16)

	if _, err := rand.Read(buffer); err != nil {
		panic(err)
	}

	return hex.EncodeToString(buffer)
}
