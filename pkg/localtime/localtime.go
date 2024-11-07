package localtime

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

const DateTimeFormat = "2006-01-02 15:04:05" //常规类型

// LocalDate 本地日期
type LocalTime struct {
	time.Time
}

func Now() *LocalTime {
	return &LocalTime{time.Now()}
}

var ZeroLocalTime = LocalTime{time.Unix(0, 0)}

// UnmarshalJSON gin bind 反射结构体
func (t *LocalTime) UnmarshalJSON(bytes []byte) (err error) {
	t.Time, err = time.ParseInLocation(DateTimeFormat, strings.Trim(string(bytes), "\""), time.Local)
	return
}

// MarshalJSON gorm marshal 序列化结构体
func (t LocalTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", t.Time.Format(DateTimeFormat))), nil
}

// Value LocalDate 转 time
func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return &zeroTime, nil
	}
	return t.Time, nil
}

// Scan gorm Scan 扫描时的数据赋值
func (t *LocalTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = LocalTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

// int64 转 时间
func TimeUnixToTime(timeUnix int64) (rtnTime *LocalTime) {
	formatTimeStr := time.Unix(timeUnix/1000, 0).Format(DateTimeFormat)
	timeRtn, _ := time.Parse(DateTimeFormat, formatTimeStr)
	rtnTime = &LocalTime{
		Time: timeRtn,
	}
	return
}

// int64 转 时间
func TimeUnixToTimeSubtract8(timeUnix int64) (rtnTime *LocalTime) {
	formatTimeStr := time.Unix(timeUnix/1000-28800, 0).Format(DateTimeFormat)
	timeRtn, _ := time.Parse(DateTimeFormat, formatTimeStr)
	rtnTime = &LocalTime{
		Time: timeRtn,
	}
	return
}

// 字符串日期 转换int64时间戳 2021-06-08 00:00:00
func TimeStringToTimeUnix(timestr string) (timeUnix int64) {
	timertn, _ := time.ParseInLocation(DateTimeFormat, timestr, time.Local) //使用parseInLocation将字符串格式化返回本地时区时间
	return timertn.UnixNano() / 1e6
}

// 字符串转time
func TimeStringToTime(timestr string) (timertn time.Time, err error) {
	timertn, err = time.Parse(DateTimeFormat, timestr)
	return
}

func TimeUnixToTimeUnix(timeUnix int64) (timestr string) {
	if timeUnix == 0 {
		return ""
	}
	//毫秒换算为秒
	if timeUnix > 9999999999 {
		timeUnix = timeUnix / 1000
	}
	return time.Unix(timeUnix, 0).Format(DateTimeFormat)
}

// 字符串日期 转换int64时间戳 2021-06-08 00:00:00
func TimeRFC3339ToTimeUnix(timestr string) (timeUnix int64) {
	timertn, _ := time.ParseInLocation(time.RFC3339, timestr, time.Local) //使用parseInLocation将字符串格式化返回本地时区时间
	return timertn.UnixMilli()
}
