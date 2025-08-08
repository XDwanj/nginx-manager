package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd 代表基础命令
var rootCmd = &cobra.Command{
	Use:   "nginx-manager",
	Short: "Nginx 配置文件管理工具",
	Long: `Nginx 配置文件管理工具可以根据 JSON 格式的配置文件自动生成 Nginx 配置文件。
支持递归扫描目录查找 nm-*.json 文件，并生成对应的 .conf 配置文件。`,
}

// Execute 执行根命令
func Execute() error {
	return rootCmd.Execute()
}
