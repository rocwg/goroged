package main

import (
	"log"

	edgeconfig "github.com/rocwg/goro-edge/runtime/config"
)

func main() {

	cfg := edgeconfig.Load()

	app, err := NewApplication(cfg)
	if err != nil {
		log.Fatalf(
			"❌ 创建 Edge Application 失败: %v",
			err,
		)
	}

	defer app.Close()

	if err := app.Run(); err != nil {
		log.Fatalf(
			"❌ Edge Application 运行失败: %v",
			err,
		)
	}
}
