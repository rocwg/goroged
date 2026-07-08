package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	// 🛠️ 注意：请根据你实际的包路径和 buf 生成的包名进行微调
	// 这里假设你生成的包名在 gen 目录下
	dictv1 "github.com/rocwg/grpc-contracts/gen/go/dictarea/v1"
	hellov1 "github.com/rocwg/grpc-contracts/gen/go/hello/v1"
)

func main() {
	// 1. 初始化 Gin 引擎（生产环境下可切换为 gin.ReleaseMode）
	r := gin.Default()

	// 2. 建立与底层 gRPC 微服务的“长连接矩阵”
	// 连接 Go 的 goro-dict-area-service (假设运行在本地或物理机 114)
	dictConn, err := grpc.NewClient("192.168.1.114:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("❌ 无法连接 Go 字典微服务: %v", err)
	}
	defer dictConn.Close()
	dictClient := dictv1.NewDictAreaServiceClient(dictConn)

	// 连接 Java 的 jaro-iam-service (假设运行在本地 9090 端口)
	javaConn, err := grpc.NewClient("127.0.0.1:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("❌ 无法连接 Java IAM 微服务: %v", err)
	}
	defer javaConn.Close()
	javaClient := hellov1.NewHelloServiceClient(javaConn)

	log.Println("🚀 [goro-http-adapter] 成功架起跨语言 gRPC 客户端连接连接矩阵！")

	// 3. 编写协议转换与数据聚合路由 (Adapter + BFF 职责)
	r.GET("/api/v1/adapter/test", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		// 【BFF 的高并发尊严】：利用 Goroutine 或通道并发去调下游，这里作为 Demo 顺延演示链路

		// 3.1 跨网络击穿 Windows 上的 Go 字典服务
		dictResp, err := dictClient.GetAreaByParent(ctx, &dictv1.GetAreaByParentRequest{
			ParentCode: "0",
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "调用 Go 字典服务失败: " + err.Error()})
			return
		}

		// 3.2 跨协议击穿 Java 的 Spring Boot gRPC 服务
		javaResp, err := javaClient.SayHello(ctx, &hellov1.SayHelloRequest{
			Name: "Goro-Adapter",
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "调用 Java IAM 服务失败: " + err.Error()})
			return
		}

		// 3.3 数据聚合，完美吐出纯净的 JSON 供网关或前端消费
		c.JSON(http.StatusOK, gin.H{
			"code":           200,
			"msg":            "success",
			"adapter_status": "OK (2026-07-07)",
			"data": gin.H{
				"from_go_dict":  dictResp.GetList(),
				"from_java_iam": javaResp.GetMessage(),
			},
		})
	})

	// 4. 适配器启动在 HTTP 8080 端口，默默为 KrakenD 提供无杂质的 L7 弹药
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("❌ 适配器启动失败: %v", err)
	}
}
