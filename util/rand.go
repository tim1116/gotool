package util

import (
	"crypto/rand"
	"math/big"
)

// RangeRand 生成区间[-m, n]的 int64安全随机数
func RangeRand(min, max int64) int64 {
	if min > max {
		panic("the min is greater than max!")
	}
	result, _ := rand.Int(rand.Reader, big.NewInt(max-min+1))
	return min + result.Int64()
}
