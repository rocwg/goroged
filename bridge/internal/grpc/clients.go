package grpc

import (
	"fmt"

	grpcgo "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	dictv1 "github.com/rocwg/grpc-contracts/gen/go/dictarea/v1"
	hellov1 "github.com/rocwg/grpc-contracts/gen/go/hello/v1"
)

const (
	dictAreaAddr = "192.168.1.114:50051"
	helloAddr    = "127.0.0.1:9090"
)

// Clients 保存 Bridge 所有 gRPC Client。
//
// 注意：
// Bridge 只负责协议转换。
// 不允许写业务逻辑。
type Clients struct {
	dictConn  *grpcgo.ClientConn
	helloConn *grpcgo.ClientConn

	DictArea dictv1.DictAreaServiceClient
	Hello    hellov1.HelloServiceClient
}

// NewClients 初始化所有 gRPC Client。
func NewClients() (*Clients, error) {

	dictConn, err := grpcgo.NewClient(
		dictAreaAddr,
		grpcgo.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("connect dictarea grpc: %w", err)
	}

	helloConn, err := grpcgo.NewClient(
		helloAddr,
		grpcgo.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		_ = dictConn.Close()
		return nil, fmt.Errorf("connect hello grpc: %w", err)
	}

	return &Clients{
		dictConn:  dictConn,
		helloConn: helloConn,

		DictArea: dictv1.NewDictAreaServiceClient(dictConn),
		Hello:    hellov1.NewHelloServiceClient(helloConn),
	}, nil
}

// Close 释放所有连接。
func (c *Clients) Close() {

	if c.dictConn != nil {
		_ = c.dictConn.Close()
	}

	if c.helloConn != nil {
		_ = c.helloConn.Close()
	}
}
