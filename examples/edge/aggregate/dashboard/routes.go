package dashboard

import (
	"github.com/gin-gonic/gin"

	edgeclients "github.com/rocwg/ged/examples/edge/clients"
)

// Register 注册 Dashboard API。
func Register(
	router *gin.RouterGroup,
	clients *edgeclients.Clients,
) {
	handler := NewHandler(clients)

	router.GET("/overview", handler.Overview)
}
