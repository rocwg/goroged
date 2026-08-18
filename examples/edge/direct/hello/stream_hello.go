package hello

import (
	"context"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"

	hellov1 "github.com/rocwg/grpc-contracts/gen/go/hello/v1"
)

// StreamHello
//
// HTTP(SSE)
//
//	↓
//
// Bridge
//
//	↓
//
// gRPC Server Streaming
func (a *Adapter) StreamHello(c *gin.Context) {

	var req StreamHelloRequest

	// 绑定 Query 参数，例如：?name=roc
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "参数格式错误",
		})
		return
	}

	ctx, cancel := context.WithTimeout(
		c.Request.Context(),
		a.streamTimeout,
	)
	defer cancel()

	stream, err := a.client.StreamHello(
		ctx,
		&hellov1.StreamHelloRequest{
			Name: req.Name,
		},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 告诉浏览器：这是一个 SSE 持续流
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")

	// Gin 支持 SSE Flush
	flusher, ok := c.Writer.(http.Flusher)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "stream unsupported",
		})
		return
	}

	// 持续接收 gRPC Stream
	for {

		resp, err := stream.Recv()

		// 服务端正常结束
		if err == io.EOF {
			return
		}

		// Stream 异常
		if err != nil {
			c.SSEvent("error", err.Error())
			flusher.Flush()
			return
		}

		// gRPC Stream -> HTTP SSE
		c.SSEvent("message", resp)
		flusher.Flush()
	}
}
