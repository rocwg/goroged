package iam

import (
	"net/http"

	"github.com/gin-gonic/gin"
	iamv1 "github.com/rocwg/grpc-contracts/gen/go/iam/v1"
)

func (a *Adapter) Login(c *gin.Context) {
	var req iamv1.GetUserRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "参数格式错误",
		})
		return
	}
	// TODO ...
}
