package utils

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/luantao/IM-base/pkg/config"
	"github.com/luantao/IM-base/pkg/json"
	"io"
	"strings"
	"time"
)

type HTTPClient struct {
	*resty.Client
	context context.Context
}

// NewHTTP is return new client
func NewHTTP(ctx context.Context) *HTTPClient {
	client := new(HTTPClient)
	client.Client = resty.New()
	client.context = ctx
	return client
}

var logName = "utils.yg_http"

func (c *HTTPClient) GetWithNoAuth(url string, responseBody interface{}) (err error) {

	// 超时3s重试3次
	var resp *resty.Response
	resp, err = c.SetRetryCount(3).SetTimeout(3 * time.Second).SetRetryMaxWaitTime(3 * time.Millisecond).R().Get(url)

	if err != nil {
		err = errors.New("网络请求错误")
		return
	}
	err = json.Unmarshal(resp.Body(), &responseBody)
	if err != nil {
		err = errors.New("网络请求错误")
		return
	}
	return err
}

// GET get 请求
func (c *HTTPClient) GetWithUrlParam(url string, jsonParams interface{}) (responseBody Response) {

	headers := make(map[string]string)
	params := make(map[string]string)
	jsonObj, _ := Struct2Map(jsonParams)
	for k, v := range jsonObj {
		params[SnakeString(k)] = fmt.Sprintf("%v", v)
	}
	// 超时3s重试3次
	resp, err := c.SetRetryCount(3).SetTimeout(3 * time.Second).SetRetryMaxWaitTime(3 * time.Millisecond).R().SetHeaders(headers).SetQueryParams(params).Get(url)
	if err != nil {
		err = errors.New("网络请求错误")
		return
	}

	err = json.Unmarshal(resp.Body(), &responseBody)
	if err != nil {
		err = errors.New("网络请求错误")
		return
	}
	return
}

// GET get 请求
func (c *HTTPClient) GetWithMap(url string, jsonParams map[string]interface{}) (body []byte, err error) {

	params := make(map[string]string)
	for k, v := range jsonParams {
		params[SnakeString(k)] = fmt.Sprintf("%v", v)
	}
	// 超时3s重试3次
	var resp *resty.Response
	resp, err = c.SetRetryCount(3).SetTimeout(3 * time.Second).SetRetryMaxWaitTime(3 * time.Millisecond).R().SetQueryParams(params).Get(url)
	if err != nil {
		err = errors.New("网络请求错误")
		return
	}
	return resp.Body(), err
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// example: http://host:port/uri/?param1=1&param2=2
func (c *HTTPClient) Get(url string, jsonParams interface{}) (responseBody Response) {

	params, err := json.Marshal(jsonParams)
	// 超时3s重试3次

	resp, err := c.SetRetryCount(3).SetTimeout(3 * time.Second).SetRetryMaxWaitTime(3 * time.Millisecond).R().SetBody(params).Get(url)

	if err != nil {
		err = errors.New("网络请求错误")
		return
	}
	err = json.Unmarshal(resp.Body(), &responseBody)
	if err != nil {
		err = errors.New("网络请求错误")
		return
	}
	return responseBody
}

// POST post 请求
func (c *HTTPClient) Post(url string, body interface{}) (responseBody Response) {

	headers := make(map[string]string)

	headers["content-type"] = "application/json; charset=utf-8" // content-type: application/json; charset=utf-8

	//params, _ := json.Marshal(jsonParams)
	// 超时3s重试3次
	resp, err := c.SetRetryCount(3).SetTimeout(3 * time.Second).SetRetryMaxWaitTime(3 * time.Millisecond).SetHeaders(headers).R().SetBody(body).Post(url)
	if err != nil {
		err = errors.New("网络请求错误")
		return
	}
	err = json.Unmarshal(resp.Body(), &responseBody)
	if err != nil {
		err = errors.New("网络请求错误")
		return
	}
	return responseBody
}

// Delete
func (c *HTTPClient) Delete(url string, body interface{}) (responseBody Response) {
	headers := make(map[string]string)

	headers["content-type"] = "application/json; charset=utf-8" // content-type: application/json; charset=utf-8

	//params, _ := json.Marshal(jsonParams)
	// 超时3s重试3次
	resp, err := c.SetRetryCount(3).SetTimeout(3 * time.Second).SetRetryMaxWaitTime(3 * time.Millisecond).SetHeaders(headers).R().SetBody(body).Delete(url)
	if err != nil {
		err = errors.New("网络请求错误")
		return
	}
	err = json.Unmarshal(resp.Body(), &responseBody)
	if err != nil {
		err = errors.New("网络请求错误")
		return
	}
	return responseBody
}

// POST post 请求
func (c *HTTPClient) Put(url string, body interface{}) (responseBody Response) {

	headers := make(map[string]string)
	headers["content-type"] = "application/json; charset=utf-8" // content-type: application/json; charset=utf-8

	//params, _ := json.Marshal(jsonParams)
	// 超时3s重试3次
	resp, err := c.SetRetryCount(3).SetTimeout(3 * time.Second).SetRetryMaxWaitTime(3 * time.Millisecond).SetHeaders(headers).R().SetBody(body).Put(url)
	if err != nil {
		err = errors.New("网络请求错误")
		return
	}
	err = json.Unmarshal(resp.Body(), &responseBody)
	if err != nil {
		err = errors.New("网络请求错误")
		return
	}
	return responseBody
}

// PATCH 请求
func (c *HTTPClient) Patch(url string, body interface{}) (responseBody Response) {
	headers := make(map[string]string)
	headers["content-type"] = "application/json; charset=utf-8" // content-type: application/json; charset=utf-8
	//params, _ :=json.Marshal(jsonParams)
	// 超时3s重试3次
	resp, err := c.SetRetryCount(3).SetTimeout(3 * time.Second).SetRetryMaxWaitTime(3 * time.Millisecond).SetHeaders(headers).R().SetBody(body).Patch(url)
	if err != nil {
		err = errors.New("网络请求错误")
		return
	}
	err = json.Unmarshal(resp.Body(), &responseBody)
	if err != nil {
		err = errors.New("网络请求错误")
		return
	}
	return responseBody
}

func (c *HTTPClient) PatchWithHeader(url string, body interface{}, headers map[string]string) (responseBody Response) {

	//headers["content-type"] = "application/json; charset=utf-8" // content-type: application/json; charset=utf-8
	//params, _ :=json.Marshal(jsonParams)
	// 超时3s重试3次
	resp, err := c.SetRetryCount(3).SetTimeout(3 * time.Second).SetRetryMaxWaitTime(3 * time.Millisecond).SetHeaders(headers).R().SetBody(body).Patch(url)
	if err != nil {
		err = errors.New("网络请求错误")
		return
	}
	err = json.Unmarshal(resp.Body(), &responseBody)
	if err != nil {
		err = errors.New("网络请求错误")
		return
	}
	return responseBody
}

func (c *HTTPClient) PATCH(url string, params interface{}, responseBody interface{}) error {

	headers := make(map[string]string)
	// 超时3s重试3次
	resp, err := c.SetRetryCount(3).SetTimeout(3 * time.Second).SetRetryMaxWaitTime(3 * time.Millisecond).SetHeaders(headers).R().SetBody(params).Patch(url)
	if err != nil {
		err = errors.New("网络请求错误")
		return err
	}
	_ = json.Unmarshal(resp.Body(), &responseBody)
	return err
}

func (c *HTTPClient) GET(url string, params map[string]string, responseBody interface{}) (err error) {

	headers := make(map[string]string)

	// 超时3s重试3次
	resp, err := c.SetTimeout(3 * time.Second).SetHeaders(headers).R().SetQueryParams(params).Get(url)
	if err != nil {
		err = errors.New("网络请求错误")
		return
	}

	err = json.Unmarshal(resp.Body(), responseBody)
	if err != nil {
		err = errors.New("网络请求错误")
		return
	}
	return
}

// POST post 请求

func (c *HTTPClient) POSTBody(url string, params interface{}, headers map[string]string) (body []byte, err error) {
	// 超时3s重试3次
	var resp *resty.Response
	resp, err = c.SetRetryCount(3).SetTimeout(3 * time.Second).SetRetryMaxWaitTime(3 * time.Millisecond).SetHeaders(headers).R().SetBody(params).Post(url)
	if err != nil {
		err = errors.New("网络请求错误")
		return
	}
	body = resp.Body()
	return
}

// POST post 请求
func (c *HTTPClient) POST(url string, jsonParams interface{}) (responseBody Response) {
	headers := make(map[string]string)
	params, _ := json.Marshal(jsonParams)
	// 超时3s重试3次
	resp, err := c.SetRetryCount(3).SetTimeout(3 * time.Second).SetRetryMaxWaitTime(3 * time.Millisecond).SetHeaders(headers).R().SetBody(params).Post(url)
	if err != nil {
		err = errors.New("网络请求错误")
		return
	}
	err = json.Unmarshal(resp.Body(), &responseBody)
	if err != nil {
		err = errors.New("网络请求错误")
		return
	}
	return responseBody
}

func (c *HTTPClient) POSTFormData(url string, params map[string]string) (responseBody Response) {
	headers := make(map[string]string, 0)
	headers["Content-Type"] = "application/json"

	// 超时3s重试3次
	resp, err := c.SetRetryCount(3).SetTimeout(3 * time.Second).SetRetryMaxWaitTime(3 * time.Millisecond).SetHeaders(headers).R().SetFormData(params).Post(url)
	if err != nil {
		err = errors.New("网络请求错误")
		return
	}
	err = json.Unmarshal(resp.Body(), &responseBody)
	if err != nil {
		err = errors.New("网络请求错误")
		return
	}
	return responseBody
}

// POST JSON
func (c *HTTPClient) PostJson(url string, requestBody interface{}) (responseBody string, err error) {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	var resp *resty.Response
	resp, err = c.SetRetryCount(3).SetTimeout(3 * time.Second).SetRetryMaxWaitTime(3 * time.Millisecond).SetHeaders(headers).R().SetBody(requestBody).Post(url)
	if err != nil {
		err = errors.New("网络请求错误")
		return
	}
	responseBody = resp.String()
	return
}

func (c *HTTPClient) NetHttpPost(url string, contentType string, body io.Reader) (responseBody []byte, err error) {
	headers := make(map[string]string)
	headers["Content-Type"] = contentType
	var resp *resty.Response
	resp, err = c.SetRetryCount(3).SetTimeout(3 * time.Second).SetRetryMaxWaitTime(3 * time.Millisecond).SetHeaders(headers).R().SetBody(body).Post(url)

	if err != nil {
		err = errors.New("网络请求错误")
		return
	}
	responseBody = resp.Body()
	return
}

// post 请求带header
func (c *HTTPClient) HttpPostWithHeader(url string, params map[string]string, header map[string]string) (result []byte, err error) {
	client := resty.New()
	var resp *resty.Response
	resp, err = client.SetRetryCount(3).SetTimeout(3 * time.Second).SetRetryMaxWaitTime(3).R().SetQueryParams(params).
		SetHeaders(header).Post(url)
	result = resp.Body()
	return
}

// GetpermissionProxyUrl
func GetpermissionProxyUrl(uri string) string {
	return fmt.Sprint(config.GetString("permission.host"), strings.Replace(uri, "admin", "permission", 1))
}

// GetWorkorderProxyUrl
func GetWorkorderProxyUrl(uri string) string {
	return fmt.Sprint(config.GetString("workorder.host"), strings.Replace(uri, "admin", "workorder", 1))
}

// GetPermissionProxyUrl
func GetPermissionProxyUrl(uri string) string {
	return fmt.Sprint(config.GetString("permission.host"), strings.Replace(uri, "admin", "permission", 1))
}

// GetHeatMapProxyUrl
func GetHeatMapProxyUrl(uri string) string {
	return fmt.Sprint(config.GetString("heatmap.host"), strings.Replace(uri, "admin", "data/heatmap", 1))
}
