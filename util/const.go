package util

import "runtime"

// 一些原有php常量的信息

// 获取当前操作系统名称
// eg:windows
func Os() string {
	return runtime.GOOS
}
