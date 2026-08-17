package hello

import (
	"time"

	hellov1 "github.com/rocwg/grpc-contracts/gen/go/hello/v1"
)

// Adapter 负责 HTTP -> gRPC 协议转换。
//
// 注意：
// Adapter 严禁编写业务逻辑。
// 这里只允许做协议转换。
type Adapter struct {
	client hellov1.HelloServiceClient

	// Unary RPC 超时时间
	unaryTimeout time.Duration

	// Server Streaming 超时时间
	streamTimeout time.Duration
}

func NewAdapter(client hellov1.HelloServiceClient) *Adapter {
	return &Adapter{
		client: client,

		// 可以以后改成配置文件读取
		unaryTimeout:  2 * time.Second,
		streamTimeout: 30 * time.Second,
	}
}
