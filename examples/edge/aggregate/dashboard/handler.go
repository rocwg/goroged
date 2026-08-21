package dashboard

import (
	"context"
	"net/http"
	"time"

	edgeclients "github.com/rocwg/ged/examples/edge/clients"
	gedhttp "github.com/rocwg/ged/http"
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
		gedhttp.WriteError(w, err)
		return
	}

	// Consumer Response
	_ = gedhttp.WriteJSON(w, http.StatusOK, resp)
}
