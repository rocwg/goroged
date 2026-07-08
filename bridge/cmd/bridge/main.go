package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	// 契约引流
	dictv1 "github.com/rocwg/grpc-contracts/gen/go/dictarea/v1"
	hellov1 "github.com/rocwg/grpc-contracts/gen/go/hello/v1"
)

func main() {
	r := gin.Default()

	// 1. 初始化 gRPC 客户端矩阵
	dictConn, err := grpc.NewClient("192.168.1.114:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("❌ 连接 Go 字典服务失败: %v", err)
	}
	defer dictConn.Close()
	dictClient := dictv1.NewDictAreaServiceClient(dictConn)

	javaConn, err := grpc.NewClient("127.0.0.1:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("❌ 连接 Java IAM 服务失败: %v", err)
	}
	defer javaConn.Close()
	javaClient := hellov1.NewHelloServiceClient(javaConn)

	log.Println("🚀 [goro-http-adapter] 纯净协议转换矩阵已就位，严禁写入聚合逻辑！")

	// ==========================================
	// 2. 纯机械的 1 对 1 协议映射（Adapter 职责）
	// ==========================================

	// 映射 2.1: 纯翻译 Go 的 GetAreaByParent 接口
	r.GET("/api/v1/dictarea/get-by-parent", func(c *gin.Context) {
		parentCode := c.DefaultQuery("parent_code", "0")

		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		resp, err := dictClient.GetAreaByParent(ctx, &dictv1.GetAreaByParentRequest{
			ParentCode: parentCode,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// 无脑透传返回
		c.JSON(http.StatusOK, resp)
	})

	// 映射 2.2: 纯翻译 Java 的 SayHello 接口
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

		resp, err := javaClient.SayHello(ctx, &hellov1.SayHelloRequest{
			Name: req.Name,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// 无脑透传返回
		c.JSON(http.StatusOK, resp)
	})

	// 适配器静静运行
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("❌ 适配器启动失败: %v", err)
	}
}
