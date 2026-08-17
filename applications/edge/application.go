package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/rocwg/goro-edge/applications/edge/aggregate/dashboard"
	edgeclients "github.com/rocwg/goro-edge/applications/edge/clients"
	"github.com/rocwg/goro-edge/applications/edge/direct/dictarea"
	"github.com/rocwg/goro-edge/applications/edge/direct/hello"
	edgeauth "github.com/rocwg/goro-edge/extensions/authentication"
	edgeobservability "github.com/rocwg/goro-edge/extensions/observability"

	edgeconfig "github.com/rocwg/goro-edge/runtime/config"
)

type Application struct {
	cfg edgeconfig.Config

	router  *gin.Engine
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

	// ---------------------------------------------------------
	// 1. HTTP Router
	// ---------------------------------------------------------

	router := gin.Default()

	// ---------------------------------------------------------
	// 2. Extensions
	// ---------------------------------------------------------

	router.Use(edgeobservability.Middleware())
	router.Use(edgeobservability.LoggingMiddleware())

	authenticator := edgeauth.NewDemoAuthenticator()
	router.Use(edgeauth.Middleware(authenticator))

	// ---------------------------------------------------------
	// 3. Provider Clients
	// ---------------------------------------------------------

	providerClients, err :=
		edgeclients.NewLoader(cfg).Load()

	if err != nil {
		return nil, fmt.Errorf(
			"load provider clients: %w",
			err,
		)
	}

	app := &Application{
		cfg:     cfg,
		router:  router,
		clients: providerClients,
	}

	// ---------------------------------------------------------
	// 4. Routes
	// ---------------------------------------------------------

	app.registerRoutes()

	return app, nil
}

// registerRoutes 注册 Edge 对外暴露的 Consumer API。
func (a *Application) registerRoutes() {

	api := a.router.Group("/api/v1")

	// ---------------------------------------------------------
	// Direct
	// ---------------------------------------------------------

	dictarea.Register(
		api.Group("/dictarea"),
		a.clients.DictArea,
	)

	hello.Register(
		api.Group("/hello"),
		a.clients.Hello,
	)

	// ---------------------------------------------------------
	// Aggregate
	// ---------------------------------------------------------

	dashboard.Register(
		api,
		a.clients,
	)
}

// Run 启动 HTTP Server。
func (a *Application) Run() error {

	log.Printf(
		"🚀 [edge] HTTP Server listening on %s",
		a.cfg.HTTP.Addr,
	)

	return a.router.Run(
		a.cfg.HTTP.Addr,
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
