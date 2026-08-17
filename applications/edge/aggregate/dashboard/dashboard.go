package dashboard

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	dictv1 "github.com/rocwg/grpc-contracts/gen/go/dictarea/v1"
	hellov1 "github.com/rocwg/grpc-contracts/gen/go/hello/v1"

	edgeclients "github.com/rocwg/goro-edge/applications/edge/clients"
)

// Handler Dashboard Aggregate Handler。
//
// Dashboard 是 Consumer API。
// 它可以调用多个 Provider，然后重新组织成
// 面向 Consumer 的响应。
type Handler struct {
	dictArea dictv1.DictAreaServiceClient
	hello    hellov1.HelloServiceClient
}

// New 创建 Dashboard Handler。
func New(
	clients *edgeclients.Clients,
) *Handler {

	return &Handler{
		dictArea: clients.DictArea,
		hello:    clients.Hello,
	}
}

// Register 注册 Dashboard API。
func Register(
	router *gin.RouterGroup,
	clients *edgeclients.Clients,
) {
	handler := New(clients)

	router.GET(
		"/dashboard/overview",
		handler.Overview,
	)
}

// Overview
//
// GET /api/v1/dashboard/overview
//
// 调用：
//
//	DictArea Provider
//	Hello Provider
//
// 然后聚合成 Consumer API Response。
func (h *Handler) Overview(
	c *gin.Context,
) {

	ctx, cancel := context.WithTimeout(
		c.Request.Context(),
		3*time.Second,
	)
	defer cancel()

	// ---------------------------------------------------------
	// 1. DictArea
	// ---------------------------------------------------------

	dictResp, err := h.dictArea.GetAreaByParent(
		ctx,
		&dictv1.GetAreaByParentRequest{
			ParentCode: "0",
		},
	)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": "调用 DictArea Provider 失败",
			},
		)

		return
	}

	// ---------------------------------------------------------
	// 2. Hello
	// ---------------------------------------------------------

	helloResp, err := h.hello.SayHello(
		ctx,
		&hellov1.SayHelloRequest{
			Name: "Goro-Edge",
		},
	)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": "调用 Hello Provider 失败",
			},
		)

		return
	}

	// ---------------------------------------------------------
	// 3. Consumer Response
	// ---------------------------------------------------------

	c.JSON(
		http.StatusOK,
		gin.H{
			"areas":   dictResp.GetList(),
			"message": helloResp.GetMessage(),
		},
	)
}
