package hello

import (
	"github.com/gin-gonic/gin"

	bridgegrpc "github.com/rocwg/goro-edge/bridge/internal/grpc"
)

func Register(r *gin.RouterGroup, clients *bridgegrpc.Clients) {

	adapter := NewAdapter(clients)

	r.GET(
		"/unary",
		adapter.SayHello,
	)

	r.GET(
		"/stream",
		adapter.StreamHello,
	)
}
