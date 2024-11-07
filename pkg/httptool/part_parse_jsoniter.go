// 非amd64类型使用（解决苹果M1机型不支持问题）
//go:build !amd64
// +build !amd64

package httptool

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
)

// 解析Response的Code、Message参数组装properties
func (hc *HTTPClient) parsePartResponse(responseBody string) (properties struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}) {
	defer func() {
		if e := recover(); e != nil {
			hc.logger.Error("http request jsoniter.parsePartResponse异常退出")
		}
	}()
	var (
		codeName    = "code"
		messageName = "message"
		codeFlg     = false
		messFlg     = false
	)
	if hc.propertiesCodeName != "" {
		codeName = hc.propertiesCodeName
	}
	if hc.propertiesMessageName != "" {
		codeName = hc.propertiesMessageName
	}
	iter := jsoniter.ParseString(jsoniter.ConfigDefault, responseBody)
	for field := iter.ReadObject(); field != ""; field = iter.ReadObject() {
		switch field {
		case codeName:
			valueType := iter.WhatIsNext()
			if valueType == jsoniter.NumberValue {
				properties.Code = fmt.Sprint(iter.ReadInt())
			} else if valueType == jsoniter.StringValue {
				properties.Code = iter.ReadString()
			} else if valueType == jsoniter.BoolValue {
				properties.Code = fmt.Sprint(iter.ReadBool())
			}
			codeFlg = true
			continue
		case messageName:
			if iter.WhatIsNext() == jsoniter.StringValue {
				properties.Message = iter.ReadString()
			}
			messFlg = true
			continue
		}
		if codeFlg && messFlg {
			break
		}
	}
	return
}
