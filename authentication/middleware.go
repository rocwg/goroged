package authentication

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	gedidentity "github.com/rocwg/ged/identity"
)

// Middleware 根据 Authorization Bearer Token 建立 Identity。
//
// HTTP:
//
//	Authorization
//	      ↓
//	Authenticator
//	      ↓
//	Identity
//	      ↓
//	context.Context
func Middleware(
	authenticator Authenticator,
) gin.HandlerFunc {

	return func(c *gin.Context) {

		authorization := c.GetHeader("Authorization")

		if authorization == "" {
			unauthorized(c, "missing authorization")
			return
		}

		const prefix = "Bearer "

		if !strings.HasPrefix(
			authorization,
			prefix,
		) {
			unauthorized(c, "invalid authorization scheme")
			return
		}

		token := strings.TrimSpace(
			strings.TrimPrefix(
				authorization,
				prefix,
			),
		)

		if token == "" {
			unauthorized(c, "empty bearer token")
			return
		}

		identity, err := authenticator.Authenticate(token)
		if err != nil {
			unauthorized(c, "invalid token")
			return
		}

		ctx := gedidentity.WithContext(
			c.Request.Context(),
			identity,
		)

		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

func unauthorized(
	c *gin.Context,
	message string,
) {
	c.AbortWithStatusJSON(
		http.StatusUnauthorized,
		gin.H{
			"code":    "UNAUTHORIZED",
			"message": message,
		},
	)
}
