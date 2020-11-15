package time

import (
	"time"
)

const TIME_LAYOUT = "2006-01-02 15:04:05"

// time.  类似php time()函数
func Time() int64 {
	return time.Now().Unix()
}

// 类似PHP strtotime()
// 暂时只支持TIME_LAYOUT格式转化 通过err判断转化是否成功
func Strtotime(layout string) (int64, error) {
	var unixTime int64
	loc, _ := time.LoadLocation("Local")
	theTime, err := time.ParseInLocation(TIME_LAYOUT, layout, loc)
	if err != nil {
		return unixTime, err
	}
	unixTime = theTime.Unix()
	return unixTime, err
}
