package hello

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	hellov1 "github.com/rocwg/grpc-contracts/gen/go/hello/v1"
)

func (a *Adapter) SayHello(c *gin.Context) {

	var req SayHelloRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "参数格式错误",
		})
		return
	}

	if req.Name == "" {
		req.Name = "World"
	}

	ctx, cancel := context.WithTimeout(
		c.Request.Context(),
		a.unaryTimeout,
	)
	defer cancel()

	resp, err := a.clients.Hello.SayHello(
		ctx,
		&hellov1.SayHelloRequest{
			Name: req.Name,
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
