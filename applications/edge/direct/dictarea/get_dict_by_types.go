package dictarea

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	dictv1 "github.com/rocwg/grpc-contracts/gen/go/dictarea/v1"
)

func (a *Adapter) GetDictByTypes(c *gin.Context) {

	var req GetDictByTypesRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "参数格式错误",
		})
		return
	}

	ctx, cancel := context.WithTimeout(
		c.Request.Context(),
		2*time.Second,
	)
	defer cancel()

	resp, err := a.client.GetDictByTypes(
		ctx,
		&dictv1.GetDictByTypesRequest{
			TypeCodes: req.TypeCodes,
		},
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}
