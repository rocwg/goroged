package dictarea

import (
	bridgegrpc "github.com/rocwg/goro-edge/bridge/internal/grpc"
)

// Adapter 负责 HTTP -> gRPC 协议转换。
//
// 注意：
// Adapter 严禁编写业务逻辑。
// 这里只允许做协议转换。
type Adapter struct {
	clients *bridgegrpc.Clients
}

func NewAdapter(clients *bridgegrpc.Clients) *Adapter {
	return &Adapter{
		clients: clients,
	}
}
