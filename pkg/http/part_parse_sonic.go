// amd64类型使用
//go:build amd64
// +build amd64

package http

import "github.com/bytedance/sonic"

// 解析Response的Code、Message参数组装properties
func (hc *HTTPClient) parsePartResponse(responseBody string) (properties struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}) {
	defer func() {
		if e := recover(); e != nil {
			hc.logger.Error("http request sonic.parsePartResponse异常退出")
		}
	}()
	var (
		codeName    = "code"
		messageName = "message"
	)
	if hc.propertiesCodeName != "" {
		codeName = hc.propertiesCodeName
	}
	if hc.propertiesMessageName != "" {
		messageName = hc.propertiesMessageName
	}

	codeRes, _ := sonic.GetFromString(responseBody, codeName)
	messRes, _ := sonic.GetFromString(responseBody, messageName)
	properties.Code, _ = codeRes.String()
	properties.Message, _ = messRes.String()
	return
}
