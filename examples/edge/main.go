package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	edgeconfig "github.com/rocwg/ged/examples/edge/config"
)

const shutdownTimeout = 5 * time.Second

func main() {

	// 1. Load Configuration
	cfg, err := edgeconfig.Load("examples/edge/config/config.json")
	if err != nil {
		log.Fatalf("❌ 加载 Edge 配置失败: %v", err)
	}

	// 2. Create Application
	app, err := NewApplication(cfg)
	if err != nil {
		log.Fatalf("create application: %v", err)
	}

	if err := runApplication(app); err != nil {
		log.Fatal(err)
	}
}

func runApplication(
	app *Application,
) error {

	stop := make(chan os.Signal, 1)

	signal.Notify(
		stop,
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer signal.Stop(stop)

	serverErr := make(chan error, 1)

	go func() {
		serverErr <- app.Run()
	}()

	select {
	case err := <-serverErr:
		if !errors.Is(err, http.ErrServerClosed) {
			return err
		}

	case <-stop:
		log.Println("🛑 [edge] shutting down")
	}

	ctx, cancel := context.WithTimeout(
		context.Background(),
		shutdownTimeout,
	)
	defer cancel()

	if err := app.Shutdown(ctx); err != nil {
		app.Close()

		return err
	}

	app.Close()

	log.Println("👋 [edge] stopped")

	return nil
}
