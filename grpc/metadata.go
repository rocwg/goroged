package grpc

import (
	"context"

	gedidentity "github.com/rocwg/ged/identity"
	gedrequest "github.com/rocwg/ged/request"
)

// PropagateMetadata 将 Edge Context 中的身份与 Request Context
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
func PropagateMetadata(ctx context.Context) context.Context {

	identity, ok := gedidentity.FromContext(ctx)
	if ok {
		ctx = gedidentity.AppendMetadata(ctx, identity)
	}

	requestContext, ok := gedrequest.FromContext(ctx)
	if ok {
		ctx = gedrequest.AppendMetadata(ctx, requestContext)
	}
	return ctx
}
