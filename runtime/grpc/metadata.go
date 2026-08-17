package grpc

import (
	"context"

	edgeidentity "github.com/rocwg/goro-edge/runtime/identity"
	edgerequest "github.com/rocwg/goro-edge/runtime/request"
)

// PropagateMetadata 将 Edge Context 中的身份与请求上下文
// 传播到 gRPC outgoing metadata。
//
// HTTP:
//
//	Request
//	   ↓
//	Edge Context
//
// gRPC:
//
//	Edge Context
//	   ↓
//	Outgoing Metadata
//	   ↓
//	Provider
//
// Runtime 只负责传播，不负责认证和业务判断。
func PropagateMetadata(ctx context.Context) context.Context {

	identity, ok := edgeidentity.FromContext(ctx)
	if ok {
		ctx = edgeidentity.AppendMetadata(
			ctx,
			identity,
		)
	}

	requestContext, ok := edgerequest.FromContext(ctx)
	if ok {
		ctx = edgerequest.AppendMetadata(
			ctx,
			requestContext,
		)
	}

	return ctx
}
