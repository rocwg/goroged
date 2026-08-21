package authentication

import (
	"encoding/json"
	"net/http"
	"strings"

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
) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(
			w http.ResponseWriter,
			r *http.Request,
		) {
			authorization := r.Header.Get("Authorization")

			if authorization == "" {
				unauthorized(w, "missing authorization")
				return
			}

			const prefix = "Bearer "

			if !strings.HasPrefix(authorization, prefix) {
				unauthorized(w, "invalid authorization scheme")
				return
			}

			token := strings.TrimSpace(
				strings.TrimPrefix(authorization, prefix),
			)

			if token == "" {
				unauthorized(w, "empty bearer token")
				return
			}

			identity, err := authenticator.Authenticate(token)
			if err != nil {
				unauthorized(w, "invalid token")
				return
			}

			ctx := gedidentity.WithContext(r.Context(), identity)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func unauthorized(
	w http.ResponseWriter,
	message string,
) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)

	_ = json.NewEncoder(w).Encode(map[string]string{
		"code":    "UNAUTHORIZED",
		"message": message,
	})
}
