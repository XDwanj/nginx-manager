package main

import (
	"log/slog"
	"os"

	"nginx-manager/cmd"
)

func main() {
	// 初始化日志
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		// AddSource: true,
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)

	// 执行命令行程序
	if err := cmd.Execute(); err != nil {
		slog.Error("程序执行失败", "error", err)
		os.Exit(1)
	}
}
