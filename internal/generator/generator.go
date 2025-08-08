package generator

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"nginx-manager/config"
)

// GenerateConfigFile 根据服务配置和默认配置生成 Nginx 配置文件
func GenerateConfigFile(serviceConfig *config.ServiceConfig, defaultConfig *config.DefaultConfig, jsonFilePath string) error {
	slog.Info("开始生成 Nginx 配置文件", "source", jsonFilePath)

	// 构造 .conf 文件路径
	confFilePath := strings.TrimSuffix(jsonFilePath, filepath.Ext(jsonFilePath)) + ".conf"

	// 生成配置内容
	content, err := generateConfigContent(serviceConfig, defaultConfig)
	if err != nil {
		slog.Error("生成配置内容失败", "error", err)
		return err
	}

	// 写入文件
	if err := os.WriteFile(confFilePath, []byte(content), 0644); err != nil {
		slog.Error("写入配置文件失败", "path", confFilePath, "error", err)
		return err
	}

	slog.Info("Nginx 配置文件生成成功", "path", confFilePath)
	return nil
}

// generateConfigContent 生成配置文件内容
func generateConfigContent(serviceConfig *config.ServiceConfig, defaultConfig *config.DefaultConfig) (string, error) {
	var sb strings.Builder

	// 写入 server 块开始
	sb.WriteString("server {\n")

	// 写入 server_name
	sb.WriteString(fmt.Sprintf("    server_name %s;\n\n", serviceConfig.ServerName))

	// 写入默认的 server 配置项
	for _, item := range defaultConfig.ServerItems {
		sb.WriteString(fmt.Sprintf("    %s\n", item))
	}

	// 写入 locations
	for i, location := range serviceConfig.Locations {
		// 写入 location 块开始
		sb.WriteString(fmt.Sprintf("    location %s {\n", location.Location))

		// 写入 proxy_pass
		sb.WriteString(fmt.Sprintf("        proxy_pass %s;\n\n", location.ProxyPass))

		// 对于第一个 location，写入默认的 location 配置项
		if i == 0 {
			for _, item := range defaultConfig.LocationFirstItems {
				sb.WriteString(fmt.Sprintf("        %s\n", item))
			}
			sb.WriteString("\n")
		}

		// 写入其他配置项
		for _, item := range location.Items {
			sb.WriteString(fmt.Sprintf("        %s\n", item))
		}

		// 写入 location 块结束
		sb.WriteString("    }\n\n")
	}

	// 写入 server 块结束
	sb.WriteString("}\n")

	return sb.String(), nil
}
