package sonyflake

import (
	"github.com/sony/sonyflake"
)

var sf *sonyflake.Sonyflake

func init() {
	var st sonyflake.Settings
	sf = sonyflake.NewSonyflake(st)
	if sf == nil {
		panic("sonyflake init panic")
	}
}

func ID() (id int64, err error) {
	uid, err := sf.NextID()
	if err != nil {
		return
	}
	id = int64(uid)
	return
}
