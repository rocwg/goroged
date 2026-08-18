package authentication_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/rocwg/ged/authentication"
	"github.com/rocwg/ged/identity"
)

var errInvalidToken = errors.New("invalid token")

type testAuthenticator struct {
	token string
}

func (a testAuthenticator) Authenticate(
	token string,
) (identity.Context, error) {

	if token != a.token {
		return identity.Context{}, errInvalidToken
	}

	return identity.Context{
		UserID:   "10001",
		TenantID: "demo",
	}, nil
}

func TestMiddleware(t *testing.T) {

	gin.SetMode(gin.TestMode)

	router := gin.New()

	router.Use(
		authentication.Middleware(
			testAuthenticator{
				token: "demo-token",
			},
		),
	)

	router.GET(
		"/test",
		func(c *gin.Context) {

			identityContext, ok :=
				identity.FromContext(
					c.Request.Context(),
				)

			if !ok {
				c.JSON(
					http.StatusInternalServerError,
					nil,
				)
				return
			}

			c.JSON(
				http.StatusOK,
				identityContext,
			)
		},
	)

	req := httptest.NewRequest(
		http.MethodGet,
		"/test",
		nil,
	)

	req.Header.Set(
		"Authorization",
		"Bearer demo-token",
	)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(
		recorder,
		req,
	)

	if recorder.Code != http.StatusOK {
		t.Fatalf(
			"got status %d, want %d",
			recorder.Code,
			http.StatusOK,
		)
	}
}

func TestMiddleware_MissingAuthorization(t *testing.T) {

	gin.SetMode(gin.TestMode)

	router := gin.New()

	router.Use(
		authentication.Middleware(
			testAuthenticator{
				token: "demo-token",
			},
		),
	)

	router.GET(
		"/test",
		func(c *gin.Context) {
			c.Status(http.StatusOK)
		},
	)

	req := httptest.NewRequest(
		http.MethodGet,
		"/test",
		nil,
	)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(
		recorder,
		req,
	)

	if recorder.Code != http.StatusUnauthorized {
		t.Fatalf(
			"got status %d, want %d",
			recorder.Code,
			http.StatusUnauthorized,
		)
	}
}

func TestMiddleware_InvalidToken(t *testing.T) {

	gin.SetMode(gin.TestMode)

	router := gin.New()

	router.Use(
		authentication.Middleware(
			testAuthenticator{
				token: "demo-token",
			},
		),
	)

	router.GET(
		"/test",
		func(c *gin.Context) {
			c.Status(http.StatusOK)
		},
	)

	req := httptest.NewRequest(
		http.MethodGet,
		"/test",
		nil,
	)

	req.Header.Set(
		"Authorization",
		"Bearer invalid-token",
	)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(
		recorder,
		req,
	)

	if recorder.Code != http.StatusUnauthorized {
		t.Fatalf(
			"got status %d, want %d",
			recorder.Code,
			http.StatusUnauthorized,
		)
	}
}
