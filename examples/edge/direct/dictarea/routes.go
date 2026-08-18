package dictarea

import (
	"github.com/gin-gonic/gin"
	dictv1 "github.com/rocwg/grpc-contracts/gen/go/dictarea/v1"
)

func Register(
	r *gin.RouterGroup,
	client dictv1.DictAreaServiceClient,
) {
	adapter := NewAdapter(client)

	r.GET("/get-by-parent", adapter.getAreaByParent)
	r.POST("/search", adapter.searchArea)
	r.POST("/batch-get-by-codes", adapter.batchGetAreaByCodes)
	r.POST("/dicts", adapter.getDictByTypes)
}
