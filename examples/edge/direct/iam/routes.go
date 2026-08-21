package iam

import (
	"net/http"

	iamv1 "github.com/rocwg/grpc-contracts/gen/go/iam/v1"
)

func Register(
	mux *http.ServeMux,
	client iamv1.AuthenticationServiceClient,
) {
	adapter := NewAdapter(client)

	mux.HandleFunc("POST /api/v1/iam/login", adapter.login)
}
