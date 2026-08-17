package dictarea

import (
	"context"

	edgeidentity "github.com/rocwg/goro-edge/runtime/identity"
	edgerequest "github.com/rocwg/goro-edge/runtime/request"
	dictv1 "github.com/rocwg/grpc-contracts/gen/go/dictarea/v1"
)

type Adapter struct {
	client dictv1.DictAreaServiceClient
}

func NewAdapter(
	client dictv1.DictAreaServiceClient,
) *Adapter {
	return &Adapter{
		client: client,
	}
}

func (a *Adapter) grpcContext(
	ctx context.Context,
) context.Context {

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
