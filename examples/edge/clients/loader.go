package clients

import (
	"fmt"

	dictv1 "github.com/rocwg/grpc-contracts/gen/go/dictarea/v1"
	hellov1 "github.com/rocwg/grpc-contracts/gen/go/hello/v1"

	edgeconfig "github.com/rocwg/ged/examples/edge/config"
	gedgrpc "github.com/rocwg/ged/grpc"
)

// Loader 负责在 Edge 启动阶段加载所有 Provider Client。
//
// Loader 只负责：
//   - 读取已经解析好的配置
//   - 创建 Provider Client
//   - 组装 Clients
//
// Loader 不负责：
//   - HTTP Route
//   - Business Logic
//   - BFF
//   - Authentication
type Loader struct {
	cfg edgeconfig.Config
}

// NewLoader 创建 Client Loader。
func NewLoader(cfg edgeconfig.Config) *Loader {
	return &Loader{
		cfg: cfg,
	}
}

// Load 加载所有 Provider Client。
func (l *Loader) Load() (*Clients, error) {

	result := &Clients{}

	// ---------------------------------------------------------
	// DictArea
	// ---------------------------------------------------------

	dictAreaConfig, ok := l.cfg.Providers["dict-area"]
	if !ok {
		return nil, fmt.Errorf(
			"provider config not found: dict-area",
		)
	}

	dictAreaConn, err := gedgrpc.NewClient(
		dictAreaConfig.Address,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"load dict-area client: %w",
			err,
		)
	}

	result.DictArea = dictv1.NewDictAreaServiceClient(
		dictAreaConn,
	)

	result.closeFuncs = append(
		result.closeFuncs,
		func() {
			gedgrpc.CloseClient(dictAreaConn)
		},
	)

	// ---------------------------------------------------------
	// Hello
	// ---------------------------------------------------------

	helloConfig, ok := l.cfg.Providers["hello"]
	if !ok {
		result.Close()

		return nil, fmt.Errorf(
			"provider config not found: hello",
		)
	}

	helloConn, err := gedgrpc.NewClient(
		helloConfig.Address,
	)
	if err != nil {
		result.Close()

		return nil, fmt.Errorf(
			"load hello client: %w",
			err,
		)
	}

	result.Hello = hellov1.NewHelloServiceClient(
		helloConn,
	)

	result.closeFuncs = append(
		result.closeFuncs,
		func() {
			gedgrpc.CloseClient(helloConn)
		},
	)

	return result, nil
}
