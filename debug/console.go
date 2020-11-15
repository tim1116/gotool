package debug

import "fmt"

// 调试使用
func Show(v interface{}) {
	fmt.Printf("%T ---> %v\n", v, v)
}
