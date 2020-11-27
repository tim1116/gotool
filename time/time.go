package time

/**
  补充时间换算规则： 秒 > 毫秒 > 微妙 > 纳秒 > ..
  1s=10^3ms(毫秒)=10^6μs(微秒)=10^9ns(纳秒)=10^12ps(皮秒)=10^15fs(飞秒)=10^18as(阿秒)=10^21zm(仄秒)=10^24ym(幺秒)
*/

import (
	"time"
)

const TIME_LAYOUT = "2006-01-02 15:04:05"

// time.  类似php time()函数
func Time() int64 {
	return time.Now().Unix()
}

// 返回当前时间微妙  有点类似php microtime()
func MicroTime() float64 {
	micro := time.Now().UnixNano() / 1e3
	time := float64(micro)
	return time / 1000000
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
