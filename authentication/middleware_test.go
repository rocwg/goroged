package authentication_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

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

	next := http.HandlerFunc(
		func(
			w http.ResponseWriter,
			r *http.Request,
		) {

			identityContext, ok :=
				identity.FromContext(
					r.Context(),
				)

			if !ok {
				http.Error(
					w,
					"identity not found",
					http.StatusInternalServerError,
				)
				return
			}

			w.WriteHeader(http.StatusOK)

			_, _ = w.Write(
				[]byte(identityContext.UserID),
			)
		},
	)

	handler := authentication.Middleware(
		testAuthenticator{
			token: "demo-token",
		},
	)(next)

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

	handler.ServeHTTP(
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

	next := http.HandlerFunc(
		func(
			w http.ResponseWriter,
			r *http.Request,
		) {
			w.WriteHeader(http.StatusOK)
		},
	)

	handler := authentication.Middleware(
		testAuthenticator{
			token: "demo-token",
		},
	)(next)

	req := httptest.NewRequest(
		http.MethodGet,
		"/test",
		nil,
	)

	recorder := httptest.NewRecorder()

	handler.ServeHTTP(
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

	next := http.HandlerFunc(
		func(
			w http.ResponseWriter,
			r *http.Request,
		) {
			w.WriteHeader(http.StatusOK)
		},
	)

	handler := authentication.Middleware(
		testAuthenticator{
			token: "demo-token",
		},
	)(next)

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

	handler.ServeHTTP(
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
