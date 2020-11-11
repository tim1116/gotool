package time

import (
	"fmt"
	"testing"
)

func TestTime(t *testing.T) {
	fmt.Println("TestTime")
	// 自己看控制台输出有无异常
	r := Time()
	fmt.Println(r)

}
