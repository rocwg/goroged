package dashboard

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	edgeclients "github.com/rocwg/ged/examples/edge/clients"
)

type Handler struct {
	service *Service
}

// NewHandler 创建 Dashboard Handler。
func NewHandler(
	clients *edgeclients.Clients,
) *Handler {
	return &Handler{
		service: NewService(clients),
	}
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
func (h *Handler) Overview(c *gin.Context) {

	ctx, cancel := context.WithTimeout(
		c.Request.Context(),
		3*time.Second,
	)
	defer cancel()

	resp, err := h.service.overview(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "获取 Dashboard Overview 失败",
		})
		return
	}

	// Consumer Response
	c.JSON(http.StatusOK, resp)
}
