package mhash

import (
	"crypto/md5"
	"fmt"
)

// Times33 MD5
func Times33(s string) string {
	hashValue := uint32(5381)
	for _, char := range s {
		hashValue = ((hashValue << 5) + hashValue) + uint32(char)
	}
	bytes := md5.Sum([]byte(s))
	return fmt.Sprintf("%x%d", bytes, hashValue)
}
