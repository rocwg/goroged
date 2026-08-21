package dashboard

import (
	"net/http"

	edgeclients "github.com/rocwg/ged/examples/edge/clients"
)

// Register 注册 Dashboard API。
func Register(
	mux *http.ServeMux,
	clients *edgeclients.Clients,
) {
	handler := NewHandler(clients)

	mux.HandleFunc("GET /api/v1/dashboard/overview", handler.overview)
}
