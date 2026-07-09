package dictarea

import (
	"github.com/gin-gonic/gin"

	bridgegrpc "github.com/rocwg/goro-edge/bridge/internal/grpc"
)

func Register(r *gin.RouterGroup, clients *bridgegrpc.Clients) {

	adapter := NewAdapter(clients)

	r.GET(
		"/get-by-parent",
		adapter.GetAreaByParent,
	)

	r.POST(
		"/search",
		adapter.SearchArea,
	)

	r.POST(
		"/batch-get-by-codes",
		adapter.BatchGetAreaByCodes,
	)

	r.POST(
		"/dicts",
		adapter.GetDictByTypes,
	)
}
