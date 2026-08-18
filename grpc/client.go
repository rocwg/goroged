package grpc

import (
	"context"
	"fmt"

	grpcgo "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// NewClient 创建 Edge → Provider gRPC ClientConn。
//
// Runtime 负责：
//   - gRPC connection
//   - Edge Context → outgoing metadata
//
// Application 不需要关心 metadata 如何传播。
func NewClient(addr string) (*grpcgo.ClientConn, error) {

	conn, err := grpcgo.NewClient(
		addr,
		grpcgo.WithTransportCredentials(
			insecure.NewCredentials(),
		),
		grpcgo.WithChainUnaryInterceptor(
			UnaryMetadataInterceptor(),
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

// UnaryMetadataInterceptor 将 Edge Context 中的上下文信息
// 自动传播到 Provider gRPC metadata。
func UnaryMetadataInterceptor() grpcgo.UnaryClientInterceptor {

	return func(
		ctx context.Context,
		method string,
		req interface{},
		reply interface{},
		cc *grpcgo.ClientConn,
		invoker grpcgo.UnaryInvoker,
		opts ...grpcgo.CallOption,
	) error {

		ctx = PropagateMetadata(ctx)

		return invoker(
			ctx,
			method,
			req,
			reply,
			cc,
			opts...,
		)
	}
}
