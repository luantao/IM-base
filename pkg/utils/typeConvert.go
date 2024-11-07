package utils

import (
	"strconv"
)

// 类型转换相关的操作
// 字符串转换为int
func String2int(s string) int {
	out, _ := strconv.Atoi(s)
	return out
}

// int到string
func Int2String(i int) string {
	out := strconv.Itoa(i)
	return out
}
