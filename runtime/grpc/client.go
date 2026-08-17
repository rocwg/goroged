package grpc

import (
	"fmt"

	grpcgo "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// NewClient 创建一个 gRPC ClientConn。
//
// runtime/grpc 只负责 gRPC 连接能力。
// 它不知道具体连接的是哪个 Provider。
func NewClient(addr string) (*grpcgo.ClientConn, error) {

	conn, err := grpcgo.NewClient(
		addr,
		grpcgo.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)

	if err != nil {
		return nil, fmt.Errorf(
			"connect grpc server %s: %w",
			addr,
			err,
		)
	}

	return conn, nil
}

// CloseClient 关闭 gRPC ClientConn。
func CloseClient(conn *grpcgo.ClientConn) {
	if conn == nil {
		return
	}

	_ = conn.Close()
}
