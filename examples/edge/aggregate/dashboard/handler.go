package dashboard

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

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

// overview
//
// GET /api/v1/dashboard/overview
//
// 调用：
//
//	DictArea Provider
//	Hello Provider
//
// 然后聚合成 Consumer API Response。
func (h *Handler) overview(
	w http.ResponseWriter,
	r *http.Request,
) {

	ctx, cancel := context.WithTimeout(
		r.Context(),
		3*time.Second,
	)
	defer cancel()

	resp, err := h.service.overview(ctx)
	if err != nil {
		http.Error(w, "获取 Dashboard Overview 失败", http.StatusInternalServerError)
		return
	}

	// Consumer Response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}
