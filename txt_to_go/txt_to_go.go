package main

import (
	"os"
	"path/filepath"
	"strings"
)

func visit(path string, f os.FileInfo, err error) error {
	if !f.IsDir() && strings.HasSuffix(path, ".txt") && !strings.HasSuffix(path, "README.txt") {
		newPath := strings.TrimSuffix(path, filepath.Ext(path)) + ".go"
		os.Rename(path, newPath)
	}
	return nil
}

func main() {
	root := "." // 基準フォルダパスを指定してください
	filepath.Walk(root, visit)
}
