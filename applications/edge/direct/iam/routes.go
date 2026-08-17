package iam

import (
	"github.com/gin-gonic/gin"
	iamv1 "github.com/rocwg/grpc-contracts/gen/go/iam/v1"
)

func Register(
	r *gin.RouterGroup,
	client iamv1.AuthenticationServiceClient,
) {
	adapter := NewAdapter(client)

	r.POST("/login", adapter.Login)
}
