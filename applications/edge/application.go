package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	edgeauth "github.com/rocwg/goro-edge/extensions/authentication"
	edgeobservability "github.com/rocwg/goro-edge/extensions/observability"

	edgeconfig "github.com/rocwg/goro-edge/runtime/config"

	edgeclients "github.com/rocwg/goro-edge/applications/edge/clients"

	"github.com/rocwg/goro-edge/applications/edge/aggregate/dashboard"
	"github.com/rocwg/goro-edge/applications/edge/direct/dictarea"
	"github.com/rocwg/goro-edge/applications/edge/direct/hello"
)

// Application 是 goro-edge 的 Composition Root。
//
// 它负责把：
//
//   - Runtime
//   - Extensions
//   - Provider Clients
//   - Direct APIs
//   - Aggregate APIs
//
// 组装成一个完整的 Edge Application。
type Application struct {
	cfg edgeconfig.Config

	router  *gin.Engine
	clients *edgeclients.Clients
}

// NewApplication 创建 Edge Application。
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

	router.Use(
		edgeobservability.Middleware(),
	)

	router.Use(
		edgeobservability.LoggingMiddleware(),
	)

	authenticator := edgeauth.NewDemoAuthenticator()

	router.Use(
		edgeauth.Middleware(authenticator),
	)

	// ---------------------------------------------------------
	// 3. Provider Clients
	// ---------------------------------------------------------

	providerClients, err := edgeclients.New(
		cfg.Provider,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"initialize provider clients: %w",
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

	if err := app.registerRoutes(); err != nil {
		providerClients.Close()

		return nil, err
	}

	return app, nil
}

// registerRoutes 注册 Consumer API。
func (a *Application) registerRoutes() error {

	api := a.router.Group("/api/v1")

	// ---------------------------------------------------------
	// Direct APIs
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
	// Aggregate APIs
	// ---------------------------------------------------------

	dashboard.Register(
		api,
		a.clients,
	)

	return nil
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

// Close 释放 Application 资源。
func (a *Application) Close() {

	if a == nil {
		return
	}

	if a.clients != nil {
		a.clients.Close()
	}
}
