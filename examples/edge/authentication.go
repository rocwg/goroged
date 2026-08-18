package main

import (
	"errors"

	gedidentity "github.com/rocwg/ged/identity"
)

var errInvalidToken = errors.New("invalid token")

type DemoAuthenticator struct{}

func NewDemoAuthenticator() *DemoAuthenticator {
	return &DemoAuthenticator{}
}

func (a *DemoAuthenticator) Authenticate(
	token string,
) (gedidentity.Context, error) {

	if token != "demo-token" {
		return gedidentity.Context{}, errInvalidToken
	}

	return gedidentity.Context{
		UserID:   "10001",
		TenantID: "demo",
	}, nil
}
