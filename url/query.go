package url

import (
	"net/url"
)

func HttpBuildQuery(params map[string]string) string {
	q := url.Values{}
	for k, v := range params {
		q.Add(k, v)
	}
	return q.Encode()
}
