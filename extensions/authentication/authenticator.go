package authentication

import (
	"errors"

	edgeidentity "github.com/rocwg/goro-edge/runtime/identity"
)

var ErrInvalidToken = errors.New("invalid token")

type Authenticator interface {
	Authenticate(token string) (edgeidentity.Context, error)
}

type DemoAuthenticator struct{}

func NewDemoAuthenticator() *DemoAuthenticator {
	return &DemoAuthenticator{}
}

func (a *DemoAuthenticator) Authenticate(
	token string,
) (edgeidentity.Context, error) {

	if token != "demo-token" {
		return edgeidentity.Context{}, ErrInvalidToken
	}

	return edgeidentity.Context{
		UserID:   "10001",
		TenantID: "demo",
	}, nil
}
