package main

import (
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
	handler http.Handler
	clients *edgeclients.Clients
}

// NewApplication 创建 Edge Application。
//
// Application 是整个 Edge 的 Composition Root：
//
//	配置
//	  ↓
//	Client Loader
//	  ↓
//	Provider Clients
//	  ↓
//	Direct / Aggregate
func NewApplication(
	cfg edgeconfig.Config,
) (*Application, error) {

	// 1. HTTP Router
	mux := http.NewServeMux()

	// 2. Provider Clients
	providerClients, err := edgeclients.NewLoader(cfg).Load()
	if err != nil {
		return nil, fmt.Errorf("load provider clients: %w", err)
	}

	// 3. Routes
	registerRoutes(mux, providerClients)

	// 4. Middleware
	authenticator := NewDemoAuthenticator()
	var handler http.Handler = mux

	handler = LoggingMiddleware()(handler)
	handler = gedauth.Middleware(authenticator)(handler)
	handler = gedrequest.Middleware()(handler)

	return &Application{
		cfg:     cfg,
		handler: handler,
		clients: providerClients,
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
func (a *Application) Run() error {

	log.Printf(
		"🚀 [edge] HTTP Server listening on %s",
		a.cfg.HTTP.Addr,
	)

	return http.ListenAndServe(
		a.cfg.HTTP.Addr,
		a.handler,
	)
}

// Close 释放 Application 持有的资源。
func (a *Application) Close() {

	if a == nil {
		return
	}

	if a.clients != nil {
		a.clients.Close()
	}
}
