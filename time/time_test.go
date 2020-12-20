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
	microTime := MicroTime()
	fmt.Printf("%f\n", microTime)
}

func TestStrtotime(t *testing.T) {
	// 自己看控制台输出 有无异常
	r, _ := Strtotime("2020-11-02 15:04:05")
	fmt.Println(r)
}

func TestDate(t *testing.T) {
	var tt int64 = 1608460259
	ss := Date(TIME_LAYOUT, tt)
	if ss != "2020-12-20 18:30:59" {
		t.FailNow()
	}
}
