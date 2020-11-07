package commonmark

import (
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	fmt.Println("start")
	r := Md5("test")
	if r != "098f6bcd4621d373cade4e832627b4f6" {
		t.Errorf("Md5() failed. Got %s, expected 098f6bcd4621d373cade4e832627b4f6.", r)
	}
}
