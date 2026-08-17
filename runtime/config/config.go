package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	HTTP      HTTPConfig                `json:"http"`
	Providers map[string]ProviderConfig `json:"providers"`
}

type HTTPConfig struct {
	Addr string `json:"addr"`
}

type ProviderConfig struct {
	Address string `json:"address"`
}

// Load 从配置文件加载 Edge 配置。
func Load(path string) (Config, error) {

	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf(
			"read config file %s: %w",
			path,
			err,
		)
	}

	var cfg Config

	if err := json.Unmarshal(data, &cfg); err != nil {
		return Config{}, fmt.Errorf(
			"decode config file %s: %w",
			path,
			err,
		)
	}

	if err := cfg.Validate(); err != nil {
		return Config{}, err
	}

	return cfg, nil
}

// Validate 校验配置。
func (c Config) Validate() error {

	if c.HTTP.Addr == "" {
		return fmt.Errorf("http.addr is required")
	}

	for name, provider := range c.Providers {
		if name == "" {
			return fmt.Errorf("provider name is required")
		}

		if provider.Address == "" {
			return fmt.Errorf(
				"provider %q address is required",
				name,
			)
		}
	}

	return nil
}
