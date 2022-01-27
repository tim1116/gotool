package slice

import (
	"fmt"
	"testing"
)

func TestCreateRandomIntSlice(t *testing.T) {
	createRandomIntSlice(5, 10)
}

func TestCreateIntSlice(t *testing.T) {
	rand := RandIntSlice{
		Min: -1,
		Max: 50,
		Len: 10,
	}
	fmt.Println(rand.createIntSlice())
}
