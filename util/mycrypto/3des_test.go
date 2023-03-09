package mycrypto

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDes3Encrypt(t *testing.T) {
	src := "hello"
	key := "5qHIk7yMbGu29SFA61234567"
	iv := "Xoji5qa9"

	data := TripleDesEncrypt(src, key, iv)
	assert.Equal(t, data, "iQv3u6hIj9c=")

	src = "user_id=6019032&phone=&expire_time=1662519025&openid=&session_key="
	key = "5qHIk7yMbGu29SFA6"
	data = TripleDesEncrypt(src, key, iv)
	assert.Equal(t, data, "U1/OIwAfq3jSQcg7rijej/D+dvemaEp5lere5c5PzgKsIb/mz1BIlVV8F82DmtpW7xHghBKCXvffjKERIRaqbOrh6UNByHBd")
}

func TestDes3ECBDecrypt(t *testing.T) {
	src := "iQv3u6hIj9c="
	key := "5qHIk7yMbGu29SFA61234567"
	iv := "Xoji5qa9"

	data, err := TripleDesDecrypt(src, key, iv)
	assert.NoError(t, err)
	t.Log(string(data))
	assert.Equal(t, data, "hello")

	// 解析异常
	src = "xxxxx"
	data, err = TripleDesDecrypt(src, key, iv)
	assert.Error(t, err)

	src = "U1/OIwAfq3jSQcg7rijej/D+dvemaEp5lere5c5PzgKsIb/mz1BIlVV8F82DmtpW7xHghBKCXvffjKERIRaqbOrh6UNByHBd"
	key = "5qHIk7yMbGu29SFA6"
	data, err = TripleDesDecrypt(src, key, iv)
	fmt.Println(data)
	assert.NoError(t, err)
	assert.Equal(t, data, "user_id=6019032&phone=&expire_time=1662519025&openid=&session_key=")
}
