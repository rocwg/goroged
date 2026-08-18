package http

import (
	"context"
	nethttp "net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const defaultTimeout = 2 * time.Second

// JSONHandler 创建一个标准的 JSON → RPC → JSON Handler。
//
// TRequest: HTTP 请求模型
// TResponse: RPC 返回模型
//
// bind:
//
//	HTTP Request → TRequest
//
// call:
//
//	TRequest → RPC → TResponse
func JSONHandler[TRequest any, TResponse any](
	bind func(*gin.Context, *TRequest) error,
	call func(context.Context, *TRequest) (*TResponse, error),
) gin.HandlerFunc {

	return func(c *gin.Context) {

		// 1. HTTP Request Binding
		var req TRequest

		if err := bind(c, &req); err != nil {
			c.JSON(nethttp.StatusBadRequest, gin.H{
				"error": "参数格式错误",
			})
			return
		}

		// 2. Request Context
		ctx, cancel := context.WithTimeout(
			c.Request.Context(),
			defaultTimeout,
		)
		defer cancel()

		// 3. RPC
		resp, err := call(ctx, &req)
		if err != nil {
			c.JSON(nethttp.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		// 4. HTTP Response
		c.JSON(nethttp.StatusOK, resp)
	}
}

//TODO gedge/http v0.2
//--------------------
//Error mapping
//Timeout configuration
//HTTP status mapping
