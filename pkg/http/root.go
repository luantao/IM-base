package http

import (
	"MyIM/pkg/config"
	"context"
	"errors"
	"net"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

var defaultConfig *Config
var defaultOnce sync.Once

// Init 初始化默认 HTTP 配置
func Init(section string, logger *Logger) {
	defaultOnce.Do(func() {
		conf, err := newConfigWithLoggerFromACM(section, logger)
		if err != nil {
			panic(err)
		}
		defaultConfig = conf
	})
}

// Default 获取默认 HTTP 配置
func Default() *Config {
	return defaultConfig
}

// DefaultClient 获取默认配置的HTTP Client
func DefaultClient(ctx context.Context) *HTTPClient {
	conf := Default()
	if conf == nil {
		return nil
	}
	return conf.Client(ctx)
}

// NewConfigWithACM 从 Config 构造配置信息
func NewConfigWithACM(section string) (conf *Config, err error) {
	buildKey := func(key string) string {
		return strings.Join([]string{section, key}, ".")
	}

	userAgent := config.GetString(buildKey("user_agent"))
	if userAgent == "" {
		err = errors.New("http UserAgent invalid")
		return
	}

	maxIdleConns := config.GetInt(buildKey("max_idle_conns"))
	if maxIdleConns <= 0 {
		maxIdleConns = DefaultMaxIdleConns
	}

	maxIdleConnsPerHost := config.GetInt(buildKey("max_idle_conns_per_host"))
	if maxIdleConnsPerHost <= 0 {
		maxIdleConnsPerHost = DefaultMaxIdleConnsPerHost
	}

	keepAliveTimeout := config.GetDuration(buildKey("keepalive_timeout")) * time.Second
	if keepAliveTimeout <= 0 {
		keepAliveTimeout = DefaultKeepAliveTimeout
	}

	idleConnTimeout := config.GetDuration(buildKey("idle_conn_timeout")) * time.Minute
	if idleConnTimeout <= 0 {
		idleConnTimeout = DefaultIdleConnTimeout
	}
	maxConnsPerHost := config.GetInt(buildKey("max_conns_per_host"))

	conf = &Config{
		UserAgent:           userAgent,
		MaxIdleConns:        maxIdleConns,
		MaxIdleConnsPerHost: maxIdleConnsPerHost,
		KeepAliveTimeout:    keepAliveTimeout,
		IdleConnTimeout:     idleConnTimeout,
		AddProperties:       config.GetBool(buildKey("add_properties")),
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				KeepAlive: keepAliveTimeout,
			}).DialContext,
			MaxIdleConnsPerHost: maxIdleConnsPerHost,
			MaxIdleConns:        maxIdleConns,
			IdleConnTimeout:     idleConnTimeout,
			MaxConnsPerHost:     maxConnsPerHost,
		},
	}

	// proxy配置
	if config.GetBool(buildKey("proxy.open")) {
		addr := config.GetString(buildKey("proxy.addr"))
		if addr == "" {
			return nil, errors.New("代理地址为空")
		}
		u, err := url.Parse(addr)
		if err != nil {
			return nil, errors.New("解析代理地址失败:" + err.Error())
		}
		transport := defaultTransport.Clone()
		transport.Proxy = http.ProxyURL(u)
		conf.Transport = transport
	}
	return
}

// newConfigWithLoggerFromACM 从 Config 构造带 logger、apm 的配置信息
func newConfigWithLoggerFromACM(section string, logger *Logger) (conf *Config, err error) {
	if logger == nil {
		err = errors.New("http logger invalid")
		return
	}

	conf, err = NewConfigWithACM(section)
	if err != nil {
		return nil, err
	}
	conf.Logger = logger
	return
}
