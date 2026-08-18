package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	gedauth "github.com/rocwg/ged/authentication"
	gedrequest "github.com/rocwg/ged/request"

	"github.com/rocwg/ged/examples/edge/aggregate/dashboard"
	edgeclients "github.com/rocwg/ged/examples/edge/clients"
	edgeconfig "github.com/rocwg/ged/examples/edge/config"
	"github.com/rocwg/ged/examples/edge/direct/dictarea"
	"github.com/rocwg/ged/examples/edge/direct/hello"
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

	router.Use(gedrequest.Middleware())
	router.Use(LoggingMiddleware())

	authenticator := NewDemoAuthenticator()
	router.Use(gedauth.Middleware(authenticator))

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

	// Direct
	hello.Register(api.Group("/hello"), a.clients.Hello)
	dictarea.Register(api.Group("/dictarea"), a.clients.DictArea)

	// Aggregate
	dashboard.Register(api.Group("/dashboard"), a.clients)
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
