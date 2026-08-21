package clients

import (
	"slices"

	dictv1 "github.com/rocwg/grpc-contracts/gen/go/dictarea/v1"
	hellov1 "github.com/rocwg/grpc-contracts/gen/go/hello/v1"
)

// Clients 保存 goro-edge 当前使用的所有 Provider Client。
//
// 注意：
// 这里不暴露 grpc.ClientConn。
// Edge 业务层只看到具体的、强类型的 RPC Client。
type Clients struct {
	DictArea dictv1.DictAreaServiceClient
	Hello    hellov1.HelloServiceClient

	closeFuncs []func()
}

// Close 关闭所有 Provider Client 对应的底层连接。
func (c *Clients) Close() {
	if c == nil {
		return
	}

	// 逆序关闭。
	for _, v := range slices.Backward(c.closeFuncs) {
		v()
	}

	c.closeFuncs = nil
}
