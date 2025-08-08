package scanner

import (
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

// FindJSONFiles 递归查找目录下所有 nm-*.json 文件
func FindJSONFiles(dirPath string) ([]string, error) {
	slog.Info("开始扫描目录查找 JSON 文件", "directory", dirPath)

	var files []string

	// 遍历目录
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			slog.Warn("访问路径时出错", "path", path, "error", err)
			return nil
		}

		// 检查是否为文件且文件名匹配 nm-*.json
		if !info.IsDir() && strings.HasPrefix(info.Name(), "nm-") && strings.HasSuffix(info.Name(), ".json") {
			files = append(files, path)
			slog.Info("找到配置文件", "path", path)
		}

		return nil
	})

	if err != nil {
		slog.Error("扫描目录时出错", "directory", dirPath, "error", err)
		return nil, err
	}

	slog.Info("目录扫描完成", "directory", dirPath, "found_files", len(files))
	return files, nil
}
