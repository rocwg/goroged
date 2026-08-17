package clients

import (
	"fmt"

	grpcgo "google.golang.org/grpc"

	dictv1 "github.com/rocwg/grpc-contracts/gen/go/dictarea/v1"
	hellov1 "github.com/rocwg/grpc-contracts/gen/go/hello/v1"

	edgeconfig "github.com/rocwg/goro-edge/runtime/config"
	edgegrpc "github.com/rocwg/goro-edge/runtime/grpc"
)

// Clients 保存 goro-edge 所使用的 Provider Clients。
//
// 注意：
//
// 这里属于 Application。
// 因此它可以知道：
//
//   - DictArea
//   - Hello
//   - IAM
//   - CMS
//
// Runtime 不应该知道这些具体业务。
type Clients struct {
	dictAreaConn *grpcgo.ClientConn
	helloConn    *grpcgo.ClientConn

	DictArea dictv1.DictAreaServiceClient
	Hello    hellov1.HelloServiceClient
}

// New 创建 Edge 所需要的 Provider Clients。
func New(cfg edgeconfig.ProviderConfig) (*Clients, error) {

	// ---------------------------------------------------------
	// DictArea
	// ---------------------------------------------------------

	dictAreaConn, err := edgegrpc.NewClient(
		cfg.DictAreaAddr,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"create dict-area client: %w",
			err,
		)
	}

	// ---------------------------------------------------------
	// Hello
	// ---------------------------------------------------------

	helloConn, err := edgegrpc.NewClient(
		cfg.HelloAddr,
	)
	if err != nil {
		edgegrpc.CloseClient(dictAreaConn)

		return nil, fmt.Errorf(
			"create hello client: %w",
			err,
		)
	}

	return &Clients{
		dictAreaConn: dictAreaConn,
		helloConn:    helloConn,

		DictArea: dictv1.NewDictAreaServiceClient(
			dictAreaConn,
		),

		Hello: hellov1.NewHelloServiceClient(
			helloConn,
		),
	}, nil
}

// Close 关闭所有 Provider Client Connection。
func (c *Clients) Close() {
	if c == nil {
		return
	}

	edgegrpc.CloseClient(c.dictAreaConn)
	edgegrpc.CloseClient(c.helloConn)
}
