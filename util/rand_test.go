package util

import "testing"

func TestRangeRand(t *testing.T) {
	num := RangeRand(10, 10)
	if num != 10 {
		t.Error(`RangeRand 异常`)
	}
	for i := 0; i < 100; i++ {
		num = RangeRand(1, 10)
		if num < 1 {
			t.Error(`RangeRand 异常`)
		} else if num > 10 {
			t.Error(`RangeRand 异常`)
		}
	}
}
