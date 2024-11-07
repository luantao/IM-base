package times

import (
	"time"
)

// 时间相关
const (
	LayoutH    = "15"
	LayoutHM   = "15:04"
	LayoutYMD  = "2006-01-02"
	LayoutTime = "2006-01-02 15:04:05"

	LayoutUNYMD  = "20060102"   // 无分隔符
	LayoutUNYMDH = "2006010215" // 无分隔符
	MinTime      = "1900-01-01 00:00:00"
	MaXTime      = "2079-06-06 00:00:00"
)

// TimeFormat 格式时间
func TimeFormat(t time.Time, layout string) string {
	if t.IsZero() {
		return ""
	}
	if layout == "" {
		layout = LayoutTime
	}
	return t.Format(layout)
}

// TimeLocal 本地时区
var TimeLocal *time.Location

func init() {
	TimeLocal, _ = time.LoadLocation("Local")
}

// 字符串时间转成时间格式
func StringToTime(str string, layout string) (time.Time, error) {
	if str == "" {
		return time.Time{}, nil
	}
	if layout == "" {
		layout = LayoutTime
	}
	return time.ParseInLocation(layout, str, TimeLocal)
}

func MilisecondToTime(milisecound int64) time.Time {
	return time.UnixMilli(milisecound)
}
