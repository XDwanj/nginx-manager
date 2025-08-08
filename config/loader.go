package config

import (
	"encoding/json"
	"log/slog"
	"os"
)

// LoadDefaultConfig 加载默认配置文件
func LoadDefaultConfig(filePath string) (*DefaultConfig, error) {
	slog.Info("加载默认配置文件", "path", filePath)

	// 读取文件内容
	data, err := os.ReadFile(filePath)
	if err != nil {
		slog.Error("读取默认配置文件失败", "path", filePath, "error", err)
		return nil, err
	}

	// 解析 JSON
	var config DefaultConfig
	if err := json.Unmarshal(data, &config); err != nil {
		slog.Error("解析默认配置文件失败", "path", filePath, "error", err)
		return nil, err
	}

	slog.Info("默认配置文件加载成功", "path", filePath)
	return &config, nil
}

// LoadServiceConfig 加载服务配置文件
func LoadServiceConfig(filePath string) (*ServiceConfig, error) {
	slog.Info("加载服务配置文件", "path", filePath)

	// 读取文件内容
	data, err := os.ReadFile(filePath)
	if err != nil {
		slog.Error("读取服务配置文件失败", "path", filePath, "error", err)
		return nil, err
	}

	// 解析 JSON
	var config ServiceConfig
	if err := json.Unmarshal(data, &config); err != nil {
		slog.Error("解析服务配置文件失败", "path", filePath, "error", err)
		return nil, err
	}

	slog.Info("服务配置文件加载成功", "path", filePath)
	return &config, nil
}
