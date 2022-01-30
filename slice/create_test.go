package slice

import (
	"testing"
)

func TestCreateRandomIntSlice(t *testing.T) {
	lenSlice := 5
	sli := CreateRandomIntSlice(lenSlice, 10)
	if len(sli) != lenSlice {
		t.Error(`切片长度异常`)
	}
}

func TestCreateIntSlice(t *testing.T) {
	vMax, vMin, vLen := 50, -1, 10
	rand := RandSlice{
		Min: int64(vMin),
		Max: int64(vMax),
		Len: vLen,
	}
	intSlice := rand.CreateIntSlice()
	if len(intSlice) != vLen {
		t.Error(`切片长度异常`)
	}
	for _, v := range intSlice {
		if v < vMin || v > vMax {
			t.Error(`范围异常`)
		}
	}
}
