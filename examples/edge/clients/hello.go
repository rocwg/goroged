package clients

import (
	hellov1 "github.com/rocwg/grpc-contracts/gen/go/hello/v1"

	edgeconfig "github.com/rocwg/ged/examples/edge/config"
	gedgrpc "github.com/rocwg/ged/grpc"
)

// newHello 创建 Hello Provider Client。
func newHello(
	cfg edgeconfig.ProviderConfig,
) (
	hellov1.HelloServiceClient,
	func(),
	error,
) {
	conn, err := gedgrpc.NewClient(cfg.Address)
	if err != nil {
		return nil, func() {}, err
	}

	client := hellov1.NewHelloServiceClient(conn)

	closeFunc := func() {
		gedgrpc.CloseClient(conn)
	}

	return client, closeFunc, nil
}
