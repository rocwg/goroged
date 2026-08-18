package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"

	gedidentity "github.com/rocwg/ged/identity"
	gedrequest "github.com/rocwg/ged/request"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		start := time.Now()

		c.Next()

		requestContext, _ :=
			gedrequest.FromContext(c.Request.Context())

		identityContext, _ :=
			gedidentity.FromContext(c.Request.Context())

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
