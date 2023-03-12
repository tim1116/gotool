package url

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestHttpBuildQuery(t *testing.T) {
	var params = map[string]string{
		"client_code":   "aaaaa",
		"action":        "getAccessToken",
		"client_secret": "ccccc",
		"interface":     "wxapi",
		"time":          strconv.FormatInt(123, 10),
	}

	assert.Equal(t, HttpBuildQuery(params), "action=getAccessToken&client_code=aaaaa&client_secret=ccccc&interface=wxapi&time=123")
}
