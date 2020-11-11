package time

import (
	"time"
)

// time.  类似php time()函数
func Time() int64 {
	return time.Now().Unix()
}
