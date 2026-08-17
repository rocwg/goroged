package clients

import (
	hellov1 "github.com/rocwg/grpc-contracts/gen/go/hello/v1"

	edgeconfig "github.com/rocwg/goro-edge/runtime/config"
	edgegrpc "github.com/rocwg/goro-edge/runtime/grpc"
)

// newHello 创建 Hello Provider Client。
func newHello(
	cfg edgeconfig.ProviderConfig,
) (
	hellov1.HelloServiceClient,
	func(),
	error,
) {
	conn, err := edgegrpc.NewClient(cfg.Address)
	if err != nil {
		return nil, func() {}, err
	}

	client := hellov1.NewHelloServiceClient(conn)

	closeFunc := func() {
		edgegrpc.CloseClient(conn)
	}

	return client, closeFunc, nil
}
