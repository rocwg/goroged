package observability

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"

	edgeidentity "github.com/rocwg/goro-edge/runtime/identity"
	edgerequest "github.com/rocwg/goro-edge/runtime/request"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		start := time.Now()

		c.Next()

		requestContext, _ :=
			edgerequest.FromContext(c.Request.Context())

		identityContext, _ :=
			edgeidentity.FromContext(c.Request.Context())

		log.Printf(
			"request completed method=%s path=%s status=%d duration=%s request_id=%s trace_id=%s user_id=%s tenant_id=%s",
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			time.Since(start),
			requestContext.RequestID,
			requestContext.TraceID,
			identityContext.UserID,
			identityContext.TenantID,
		)
	}
}
