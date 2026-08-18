package dictarea

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	dictv1 "github.com/rocwg/grpc-contracts/gen/go/dictarea/v1"
)

func (a *Adapter) getAreaByParent(c *gin.Context) {

	// 1. Bind Request
	var req GetAreaByParentRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "参数格式错误",
		})
		return
	}

	if req.ParentCode == "" {
		req.ParentCode = "0"
	}

	// 2. context.WithTimeout()
	ctx, cancel := context.WithTimeout(
		c.Request.Context(),
		2*time.Second,
	)
	defer cancel()

	// 3. grpc Request
	// 4. grpc Call
	resp, err := a.client.GetAreaByParent(
		ctx,
		&dictv1.GetAreaByParentRequest{
			ParentCode: req.ParentCode,
		},
	)

	// 5. Return JSON
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 无脑透传返回
	c.JSON(http.StatusOK, resp)
}
