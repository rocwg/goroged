package dictarea

import (
	"time"

	dictv1 "github.com/rocwg/grpc-contracts/gen/go/dictarea/v1"
)

type Adapter struct {
	client       dictv1.DictAreaServiceClient
	unaryTimeout time.Duration
}

func NewAdapter(
	client dictv1.DictAreaServiceClient,
) *Adapter {
	return &Adapter{
		client:       client,
		unaryTimeout: 2 * time.Second,
	}
}
