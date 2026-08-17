package authentication

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	edgeidentity "github.com/rocwg/goro-edge/runtime/identity"
)

func Middleware(
	provider Authenticator,
) gin.HandlerFunc {

	return func(c *gin.Context) {

		authorization := c.GetHeader("Authorization")

		if authorization == "" {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"code":    "UNAUTHORIZED",
					"message": "missing authorization",
				},
			)
			return
		}

		const prefix = "Bearer "

		if !strings.HasPrefix(
			authorization,
			prefix,
		) {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"code":    "UNAUTHORIZED",
					"message": "invalid authorization scheme",
				},
			)
			return
		}

		token := strings.TrimSpace(
			strings.TrimPrefix(
				authorization,
				prefix,
			),
		)

		if token == "" {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"code":    "UNAUTHORIZED",
					"message": "empty bearer token",
				},
			)
			return
		}

		identity, err := provider.Authenticate(token)
		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"code":    "UNAUTHORIZED",
					"message": "invalid token",
				},
			)
			return
		}

		ctx := edgeidentity.WithContext(
			c.Request.Context(),
			identity,
		)

		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
