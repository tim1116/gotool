package slice_test

import (
	"github.com/tim1116/gotool/slice"
	"reflect"
	"testing"
)

func TestSpliceValue(t *testing.T) {
	res := slice.SpliceInsertValue([]int{1, 2, 3}, 10, 2)
	data, ok := res.([]int)
	if !ok {
		t.Error(`类型异常`)
	}
	if !reflect.DeepEqual(data, []int{1, 2, 10, 3}) {
		t.Error(`insert Error`)
	}

	res = slice.SpliceInsertValue([]int{1, 2, 3}, 10, 0)
	data, _ = res.([]int)
	if !reflect.DeepEqual(data, []int{10, 1, 2, 3}) {
		t.Error(`insert Error`)
	}

	res = slice.SpliceInsertValue([]string{"aa", "bb", "cc"}, "dd", 3)
	data2, ok := res.([]string)
	if !ok {
		t.Error(`string切片返回类型断言异常`)
	}
	if !reflect.DeepEqual(data2, []string{"aa", "bb", "cc", "dd"}) {
		t.Error(`insert Error`)
	}
}
