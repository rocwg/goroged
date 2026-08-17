package clients

import (
	"fmt"

	edgeconfig "github.com/rocwg/goro-edge/runtime/config"
)

// Loader 负责在 Edge 启动阶段加载所有 Provider Client。
//
// Loader 只负责：
//  1. 读取已经解析好的配置
//  2. 创建 Provider Client
//  3. 组装 Clients
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

	dictAreaClient, closeDictArea, err :=
		newDictArea(dictAreaConfig)

	if err != nil {
		return nil, fmt.Errorf(
			"load dict-area client: %w",
			err,
		)
	}

	result.DictArea = dictAreaClient
	result.closeFuncs = append(
		result.closeFuncs,
		closeDictArea,
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

	helloClient, closeHello, err :=
		newHello(helloConfig)

	if err != nil {
		result.Close()

		return nil, fmt.Errorf(
			"load hello client: %w",
			err,
		)
	}

	result.Hello = helloClient
	result.closeFuncs = append(
		result.closeFuncs,
		closeHello,
	)

	return result, nil
}
