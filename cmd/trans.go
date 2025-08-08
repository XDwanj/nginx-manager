package cmd

import (
	"log/slog"
	"os"
	"path/filepath"

	"nginx-manager/config"
	"nginx-manager/internal/generator"
	"nginx-manager/internal/scanner"

	"github.com/spf13/cobra"
)

// 默认配置文件路径
const defaultConfigFile = "default.json"

// transCmd 代表 trans 命令
var transCmd = &cobra.Command{
	Use:   "trans [目录路径]",
	Short: "转换 JSON 配置文件为 Nginx 配置文件",
	Long: `递归扫描指定目录，查找 nm-*.json 文件，
然后根据这些文件生成对应的 .conf 配置文件。`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dirPath := args[0]

		// 获取默认配置文件路径
		defaultPath, err := cmd.Flags().GetString("default")
		if err != nil {
			slog.Error("获取默认配置文件参数失败", "error", err)
			return
		}

		// 如果没有指定默认配置文件，则使用当前目录下的 default.json
		if defaultPath == "" {
			wd, err := os.Getwd()
			if err != nil {
				slog.Error("获取当前工作目录失败", "error", err)
				return
			}
			defaultPath = filepath.Join(wd, defaultConfigFile)
		}

		// 检查默认配置文件是否存在
		if _, err := os.Stat(defaultPath); os.IsNotExist(err) {
			slog.Error("默认配置文件不存在", "path", defaultPath)
			return
		}

		slog.Info("开始转换配置文件", "directory", dirPath, "default", defaultPath)

		// 加载默认配置
		defaultConfig, err := config.LoadDefaultConfig(defaultPath)
		if err != nil {
			slog.Error("加载默认配置失败", "error", err)
			return
		}

		// 扫描目录查找 JSON 文件
		jsonFiles, err := scanner.FindJSONFiles(dirPath)
		if err != nil {
			slog.Error("扫描目录查找 JSON 文件失败", "error", err)
			return
		}

		// 处理每个找到的 JSON 文件
		for _, jsonFile := range jsonFiles {
			// 加载服务配置
			serviceConfig, err := config.LoadServiceConfig(jsonFile)
			if err != nil {
				slog.Error("加载服务配置失败", "file", jsonFile, "error", err)
				continue
			}

			// 生成 Nginx 配置文件
			if err := generator.GenerateConfigFile(serviceConfig, defaultConfig, jsonFile); err != nil {
				slog.Error("生成 Nginx 配置文件失败", "file", jsonFile, "error", err)
				continue
			}
		}

		slog.Info("配置文件转换完成", "processed_files", len(jsonFiles))
	},
}

func init() {
	// 添加命令到根命令
	rootCmd.AddCommand(transCmd)

	// 添加命令行标志
	transCmd.Flags().StringP("default", "d", "", "指定默认配置文件路径")
}
