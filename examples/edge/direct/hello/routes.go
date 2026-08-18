package hello

import (
	"github.com/gin-gonic/gin"
	hellov1 "github.com/rocwg/grpc-contracts/gen/go/hello/v1"
)

func Register(
	r *gin.RouterGroup,
	client hellov1.HelloServiceClient,
) {
	adapter := NewAdapter(client)

	r.GET("/unary", adapter.SayHello)
	r.GET("/stream", adapter.StreamHello)
}
