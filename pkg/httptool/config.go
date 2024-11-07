package httptool

import (
	"context"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

const (
	DefaultKeepAliveTimeout    = 15 * time.Second
	DefaultIdleConnTimeout     = 5 * time.Minute
	DefaultMaxIdleConnsPerHost = 500
	DefaultMaxIdleConns        = 2000
	DefaultUserAgent           = "JRYG"
)

// Config 配置信息
//
// yaml example:
// http:
//
//	user_agent: xx-order-center
//	max_idle_conns: 2000 # 最大空闲链接数
//	keepalive_timeout: 15 # 链接保活时间;单位:秒
//	idle_conn_timeout: 5 # 空闲链接超时时间;单位:分钟
//	max_idle_conns_per_host: 500 # 单个域名最大空闲链接数
//	add_properties: true # 响应code、message添加到properties中
type Config struct {
	UserAgent           string            `json:"user_agent" yaml:"user_agent"`
	MaxIdleConns        int               `json:"max_idle_conns" yaml:"max_idle_conns"`
	KeepAliveTimeout    time.Duration     `json:"keepalive_timeout" yaml:"keepalive_timeout"`
	IdleConnTimeout     time.Duration     `json:"idle_conn_timeout" yaml:"idle_conn_timeout"`
	MaxIdleConnsPerHost int               `json:"max_idle_conns_per_host" yaml:"max_idle_conns_per_host"`
	AddProperties       bool              `json:"add_properties" yaml:"add_properties"`
	Transport           http.RoundTripper `json:"-" yaml:"-"` // 覆盖配置参数
	Logger              *Logger           `json:"-" yaml:"-"`
}

// SetTransport 修改Transport
func (c *Config) SetTransport(transport http.RoundTripper) {
	c.Transport = transport
}

func (c *Config) Client(ctx context.Context) *HTTPClient {
	if c.Transport == nil {
		c.Transport = defaultTransport
	}

	cli := resty.NewWithClient(&http.Client{
		Transport: c.Transport,
	})

	client := (&HTTPClient{
		client:        cli,
		context:       ctx,
		addProperties: c.AddProperties,
		logResponse:   true,
		logRequest:    true,
	}).Transport(c.Transport)

	if c.UserAgent != "" {
		client.UserAgent(c.UserAgent)
	}

	if c.Logger != nil {
		client.Logger(c.Logger)
	}

	return client
}
