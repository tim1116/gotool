package slice

import (
	"github.com/tim1116/gotool/util"
	"math/rand"
	"time"
)

type RandIntSlice struct {
	Min, Max int64
	Len      int
}

// RandIntSlice 返回随机int类型切片
func (r RandIntSlice) createIntSlice() []int {
	var res = make([]int, r.Len)
	for i := 0; i < r.Len; i++ {
		res[i] = int(util.RangeRand(r.Min, r.Max))
	}
	return res
}

// createRandomIntSlice 生成长度为len的int类型随机切片
// 切片中元素的最大值为max
func createRandomIntSlice(len int, max int) []int {
	var res = make([]int, len)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len; i++ {
		res[i] = rand.Intn(max)
	}
	return res
}
