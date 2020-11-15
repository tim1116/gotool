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

func TestStrtotime(t *testing.T) {
	// 自己看控制台输出有无异常
	r, _ := Strtotime("2020-11-02 15:04:05")
	fmt.Println(r)
}
