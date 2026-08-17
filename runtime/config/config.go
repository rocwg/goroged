package config

import "os"

// Config 是 goro-edge 的运行时配置。
type Config struct {
	HTTP     HTTPConfig
	Provider ProviderConfig
}

// HTTPConfig HTTP Server 配置。
type HTTPConfig struct {
	Addr string
}

// ProviderConfig Edge 所依赖的 Provider 地址。
//
// 注意：
//
// 当前 Provider 数量较少，显式列出地址。
// 暂时不使用 map[string]string，避免过度抽象。
type ProviderConfig struct {
	DictAreaAddr string
	HelloAddr    string
}

// Load 加载配置。
//
// 当前阶段：
// 环境变量 + 默认值。
func Load() Config {
	return Config{
		HTTP: HTTPConfig{
			Addr: env(
				"GORO_EDGE_HTTP_ADDR",
				":8080",
			),
		},

		Provider: ProviderConfig{
			DictAreaAddr: env(
				"GORO_EDGE_DICT_AREA_ADDR",
				"192.168.1.114:50051",
			),

			HelloAddr: env(
				"GORO_EDGE_HELLO_ADDR",
				"127.0.0.1:9090",
			),
		},
	}
}

func env(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}
