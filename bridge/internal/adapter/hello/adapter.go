package hello

import (
	"time"

	bridgegrpc "github.com/rocwg/goro-edge/bridge/internal/grpc"
)

// Adapter 负责 HTTP -> gRPC 协议转换。
//
// 注意：
// Adapter 严禁编写业务逻辑。
// 这里只允许做协议转换。
type Adapter struct {
	clients *bridgegrpc.Clients

	// Unary RPC 超时时间
	unaryTimeout time.Duration

	// Server Streaming 超时时间
	streamTimeout time.Duration
}

func NewAdapter(clients *bridgegrpc.Clients) *Adapter {
	return &Adapter{
		clients: clients,

		// 可以以后改成配置文件读取
		unaryTimeout:  2 * time.Second,
		streamTimeout: 30 * time.Second,
	}
}
