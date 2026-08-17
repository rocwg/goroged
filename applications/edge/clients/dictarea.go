package clients

import (
	dictv1 "github.com/rocwg/grpc-contracts/gen/go/dictarea/v1"

	edgeconfig "github.com/rocwg/goro-edge/runtime/config"
	edgegrpc "github.com/rocwg/goro-edge/runtime/grpc"
)

// newDictArea 创建 DictArea Provider Client。
func newDictArea(
	cfg edgeconfig.ProviderConfig,
) (
	dictv1.DictAreaServiceClient,
	func(),
	error,
) {
	conn, err := edgegrpc.NewClient(cfg.Address)
	if err != nil {
		return nil, func() {}, err
	}

	client := dictv1.NewDictAreaServiceClient(conn)

	closeFunc := func() {
		edgegrpc.CloseClient(conn)
	}

	return client, closeFunc, nil
}
