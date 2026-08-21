package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	gedauth "github.com/rocwg/ged/authentication"
	edgeclients "github.com/rocwg/ged/examples/edge/clients"
	edgeconfig "github.com/rocwg/ged/examples/edge/config"
	gedrequest "github.com/rocwg/ged/request"

	"github.com/rocwg/ged/examples/edge/aggregate/dashboard"
	"github.com/rocwg/ged/examples/edge/direct/dictarea"
	"github.com/rocwg/ged/examples/edge/direct/hello"
)

type Application struct {
	cfg     edgeconfig.Config
	server  *http.Server
	clients *edgeclients.Clients
}

// NewApplication 创建 Edge Application。
//
// Application 是整个 Edge 的 Composition Root
func NewApplication(
	cfg edgeconfig.Config,
) (*Application, error) {

	// 0. 创建 Client
	clients, err := edgeclients.NewLoader(cfg).Load()
	if err != nil {
		return nil, fmt.Errorf(
			"load provider clients: %w",
			err,
		)
	}

	// 1. 创建 Router
	mux := http.NewServeMux()

	// 2. 注册 Route
	registerRoutes(mux, clients)

	// 3. 组合 Middleware
	authenticator := NewDemoAuthenticator()
	var handler http.Handler = mux

	handler = LoggingMiddleware()(handler)
	handler = gedauth.Middleware(authenticator)(handler)
	handler = gedrequest.Middleware()(handler)

	// 4. 创建 Server
	server := &http.Server{
		Addr:    cfg.HTTP.Addr,
		Handler: handler,
	}

	return &Application{
		cfg:     cfg,
		server:  server,
		clients: clients,
	}, nil
}

// registerRoutes 注册 Edge 对外暴露的 Consumer API。
func registerRoutes(
	mux *http.ServeMux,
	clients *edgeclients.Clients,
) {
	// Direct
	hello.Register(mux, clients.Hello)
	dictarea.Register(mux, clients.DictArea)
	// Aggregate
	dashboard.Register(mux, clients)
}

// Run 启动 HTTP Server。
// ↓
// 开始服务
func (a *Application) Run() error {

	log.Printf(
		"🚀 [edge] HTTP Server listening on %s",
		a.cfg.HTTP.Addr,
	)

	return a.server.ListenAndServe()
}

// Shutdown
// ↓
// 停止 HTTP Server
// ↓
// 等待请求结束
func (a *Application) Shutdown(
	ctx context.Context,
) error {

	if a == nil || a.server == nil {
		return nil
	}

	return a.server.Shutdown(ctx)
}

// Close 释放 Application 持有的资源。
// ↓
// 释放 Provider Clients
func (a *Application) Close() {

	if a == nil {
		return
	}

	if a.clients != nil {
		a.clients.Close()
	}
}
