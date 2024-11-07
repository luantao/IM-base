package arrays

import (
	"strconv"
)

// ConverseToInt64 将字符串数组转为int64数组
func ConverseToInt64(arr []string) []int64 {
	arrInt64 := make([]int64, 0)
	for _, v := range arr {
		vv, _ := strconv.ParseInt(v, 10, 64)
		arrInt64 = append(arrInt64, vv)
	}
	return arrInt64
}

// Int64sUnique 数组去重
func Int64sUnique(array []int64) (unique []int64) {
	unique = make([]int64, 0)
	for i := 0; i < len(array); i++ {
		if Int64sContains(unique, array[i]) == -1 {
			unique = append(unique, array[i])
		}
	}
	return
}

// Int64sDifference Int64取两个数组的差集
func Int64sDifference(a, b []int64) (diff []int64) {
	for _, v := range a {
		if Int64sContains(b, v) == -1 {
			diff = append(diff, v)
		}
	}
	return
}
