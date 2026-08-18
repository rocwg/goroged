package clients

import (
	dictv1 "github.com/rocwg/grpc-contracts/gen/go/dictarea/v1"

	edgeconfig "github.com/rocwg/ged/examples/edge/config"
	gedgrpc "github.com/rocwg/ged/grpc"
)

// newDictArea 创建 DictArea Provider Client。
func newDictArea(
	cfg edgeconfig.ProviderConfig,
) (
	dictv1.DictAreaServiceClient,
	func(),
	error,
) {
	conn, err := gedgrpc.NewClient(cfg.Address)
	if err != nil {
		return nil, func() {}, err
	}

	client := dictv1.NewDictAreaServiceClient(conn)

	closeFunc := func() {
		gedgrpc.CloseClient(conn)
	}

	return client, closeFunc, nil
}
