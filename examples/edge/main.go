package main

import (
	"log"

	edgeconfig "github.com/rocwg/ged/examples/edge/config"
)

func main() {

	// ---------------------------------------------------------
	// 1. Load Configuration
	// ---------------------------------------------------------

	cfg, err := edgeconfig.Load(
		"examples/edge/config/config.json",
	)

	if err != nil {
		log.Fatalf(
			"❌ 加载 Edge 配置失败: %v",
			err,
		)
	}

	// ---------------------------------------------------------
	// 2. Create Application
	// ---------------------------------------------------------

	app, err := NewApplication(cfg)

	if err != nil {
		log.Fatalf(
			"❌ 创建 Edge Application 失败: %v",
			err,
		)
	}

	defer app.Close()

	// ---------------------------------------------------------
	// 3. Run
	// ---------------------------------------------------------

	if err := app.Run(); err != nil {
		log.Fatalf(
			"❌ Edge Application 运行失败: %v",
			err,
		)
	}
}
