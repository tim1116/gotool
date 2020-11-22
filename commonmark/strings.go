package commonmark

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// md5.
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// 只用于去除字符串前后空格使用
// 此方法和 strings.TrimSpace 一致 推荐使用TrimSpace
func Trim(str string) string {
	return strings.Trim(str, " ")
}
