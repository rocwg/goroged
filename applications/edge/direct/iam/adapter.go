package iam

import (
	iamv1 "github.com/rocwg/grpc-contracts/gen/go/iam/v1"
)

type Adapter struct {
	client iamv1.AuthenticationServiceClient
}

func NewAdapter(client iamv1.AuthenticationServiceClient) *Adapter {
	return &Adapter{
		client: client,
	}
}
