package util

import (
	"os"
	"runtime"
)

// 一些原有php常量的信息

// Os 获取当前操作系统名称
// eg:windows
func Os() string {
	return runtime.GOOS
}

// DirectorySeparator 获取文件分割符 php 中的 DIRECTORY_SEPARATOR
func DirectorySeparator() string {
	separator := os.PathSeparator
	return string(separator)
}
