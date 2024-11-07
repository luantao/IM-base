package time_utils

import (
	"fmt"
	"github.com/luantao/IM-base/pkg/times"
	"time"
)

const (
	// 定义每分钟的秒数
	SecondsPerMinute = 60
	// 定义每小时的秒数
	SecondsPerHour = SecondsPerMinute * 60
	// 定义每天的秒数
	SecondsPerDay = SecondsPerHour * 24
)

// 字符串转换为时间
func Str2Time(formatTimeStr string) time.Time {
	timeLayout := times.LayoutTime
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(timeLayout, formatTimeStr, loc) //使用模板在对应时区转化为time.time类型

	return theTime
}

// 将传入的“秒”解析为3种时间单位
func ResolveTime(seconds int) (day int, hour int, minute int) {
	day = seconds / SecondsPerDay
	hour = seconds / SecondsPerHour
	minute = seconds / SecondsPerMinute
	return
}

// TransToHour 转换为小时
func TransToHour(duration time.Duration) string {
	return fmt.Sprintf("%.0f", duration.Hours())
}

func TransMysqlTimetoString() {

}
