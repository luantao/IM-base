package id

import (
	"crypto/rand"
	"encoding/hex"
	"io"
)

type random struct {
	length uint // 长度
}

// 随机 ID 默认长度
const defaultRandomLen = 8 // 8 * 2

var rander = rand.Reader

func randomBits(b []byte) {
	if _, err := io.ReadFull(rander, b); err != nil {
		panic(err.Error())
	}
}

// RandomID 生成 16 位随机 ID.
func RandomID() string {
	return Random().ID()
}

// 随机 ID
func Random() *random {
	return &random{
		length: defaultRandomLen,
	}
}

// WithLen 指定长度
// 尽量使用偶数
func (r *random) WithLen(length uint) *random {
	r.length = length / 2
	return r
}

// ID 生成随机 ID，默认 16 位
func (r random) ID() string {
	b := make([]byte, r.length)
	randomBits(b)
	return hex.EncodeToString(b)
}
