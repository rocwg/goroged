package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rocwg/goro-edge/bridge/internal/adapter/dictarea"
	"github.com/rocwg/goro-edge/bridge/internal/adapter/hello"

	bridgegrpc "github.com/rocwg/goro-edge/bridge/internal/grpc"
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

	//
	api := r.Group("/api/v1")

	//
	dictarea.Register(api.Group("/dictarea"), clients)
	hello.Register(api.Group("/hello"), clients)

	// 运行
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("❌ Bridge 启动失败: %v", err)
	}
}
