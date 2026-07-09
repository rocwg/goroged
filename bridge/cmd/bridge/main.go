package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	bridgegrpc "github.com/rocwg/goro-edge/bridge/internal/grpc"

	dictv1 "github.com/rocwg/grpc-contracts/gen/go/dictarea/v1"
	hellov1 "github.com/rocwg/grpc-contracts/gen/go/hello/v1"
)

func main() {
	r := gin.Default()

	// 初始化 gRPC Client
	clients, err := bridgegrpc.NewClients()
	if err != nil {
		log.Fatalf("❌ 初始化 gRPC Client 失败: %v", err)
	}
	defer clients.Close()

	log.Println("🚀 [goro-bridge] 纯净协议转换矩阵已就位，严禁写入聚合逻辑！")

	// ==========================================
	// 纯机械的 1:1 协议映射
	// ==========================================

	r.GET("/api/v1/dictarea/get-by-parent", func(c *gin.Context) {
		parentCode := c.DefaultQuery("parent_code", "0")

		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		resp, err := clients.DictArea.GetAreaByParent(ctx, &dictv1.GetAreaByParentRequest{
			ParentCode: parentCode,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// 无脑透传返回
		c.JSON(http.StatusOK, resp)
	})

	r.POST("/api/v1/hello/say", func(c *gin.Context) {
		var req struct {
			Name string `json:"name"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "参数格式错误"})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		resp, err := clients.Hello.SayHello(ctx, &hellov1.SayHelloRequest{
			Name: req.Name,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// 无脑透传返回
		c.JSON(http.StatusOK, resp)
	})

	// 运行
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("❌ Bridge 启动失败: %v", err)
	}
}
