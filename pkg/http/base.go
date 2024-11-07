package http

import (
	"errors"
)

type MethodData struct {
	Config     string
	Url        string
	Timeout    int64 //单位:毫秒
	RetryCount int
	Header     map[string]string
	Limit      LimitOption
}

func (this *MethodData) validation() error {

	if len(this.Url) == 0 {
		return errors.New(this.Config + "未设置url")
	}

	if this.Timeout <= 0 {
		this.Timeout = 500
	}

	if this.RetryCount <= 0 {
		this.RetryCount = 0
	}

	if this.Header == nil {
		this.Header = make(map[string]string)
	}

	return nil
}
