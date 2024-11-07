package http

import (
	"context"
	"crypto/tls"
	"errors"
	sentinelbase "github.com/alibaba/sentinel-golang/core/base"
	"github.com/go-resty/resty/v2"
	"github.com/luantao/IM-base/pkg/id"
	xlimit "github.com/luantao/IM-base/pkg/limit"
	varctx "github.com/luantao/IM-base/pkg/var/context"
	varheader "github.com/luantao/IM-base/pkg/var/header"
	"net"
	"net/http"
	"time"
)

// defaultTransport default http client transport
var defaultTransport = &http.Transport{
	Proxy: http.ProxyFromEnvironment,
	DialContext: (&net.Dialer{
		KeepAlive: DefaultKeepAliveTimeout,
	}).DialContext,
	MaxIdleConnsPerHost: DefaultMaxIdleConnsPerHost,
	MaxIdleConns:        DefaultMaxIdleConns,
	IdleConnTimeout:     DefaultIdleConnTimeout,
}

type HTTPClient struct {
	client                *resty.Client
	logger                *Logger
	transport             http.RoundTripper
	context               context.Context
	trace                 bool // 是否跟踪链接耗时
	userAgent             string
	headers               map[string]string
	timeout               time.Duration
	retryCount            int
	uri                   string // e.g. http://dic.xx.com/city/show
	limitOpts             *limitOptions
	limitEntry            *sentinelbase.SentinelEntry
	addProperties         bool
	propertiesCodeName    string
	propertiesMessageName string
	authOption            *apiAuthOption
	logResponse           bool // 是否打response body日志
	logRequest            bool // 是否打request body,query param,form data日志
	checkHttpCode         bool // 是否校验http code，如果非2xx，则返回错误
}

func Client(ctx context.Context) *HTTPClient {
	client := resty.NewWithClient(&http.Client{
		Transport: defaultTransport,
	})

	return &HTTPClient{
		client:      client,
		context:     ctx,
		logResponse: true,
		logRequest:  true,
	}
}

// UserAgent Headers set User-Agent header
func (c *HTTPClient) UserAgent(userAgent string) *HTTPClient {
	c.userAgent = userAgent
	return c
}

// Headers set headers
func (c *HTTPClient) Headers(headers map[string]string) *HTTPClient {
	c.headers = headers
	return c
}

// Timeout set timeout
func (c *HTTPClient) Timeout(timeout time.Duration) *HTTPClient {
	c.timeout = timeout
	return c
}

// RetryCount set retry count
//
// 不设置或设置为 0,则重试 2 次
// 设置为负数则不进行重试
func (c *HTTPClient) RetryCount(count int) *HTTPClient {
	c.retryCount = count
	return c
}

// Trace 是否统计链接耗时
//
// 默认不开启,尽量通过 Config 配置
func (c *HTTPClient) Trace(trace bool) *HTTPClient {
	c.trace = trace
	return c
}

// LogResponse 日志是否打印response
// 默认开启，打印response
func (c *HTTPClient) LogResponse(logResponse bool) *HTTPClient {
	c.logResponse = logResponse
	return c
}

// LogRequest 日志是否打印request body,query param,form data
// 默认开启，打印request body,query param,form data
func (c *HTTPClient) LogRequest(logRequest bool) *HTTPClient {
	c.logRequest = logRequest
	return c
}

// CheckHttpCode 是否验证http response code，如果非2xx，则http返回error
// 默认关闭
func (c *HTTPClient) CheckHttpCode(checkHttpCode bool) *HTTPClient {
	c.checkHttpCode = checkHttpCode
	return c
}

func (c *HTTPClient) Logger(logger *Logger) *HTTPClient {
	c.logger = logger
	return c
}

func (c *HTTPClient) Transport(transport http.RoundTripper) *HTTPClient {
	c.transport = transport
	return c
}

// Headers set headers
func (c *HTTPClient) Authorization(option ApiAuthOption) *HTTPClient {
	opt := &apiAuthOption{}
	opt.applyOpts(option)
	c.authOption = opt
	return c
}

// CodeMessageName 设置code message名称
func (c *HTTPClient) CodeMessageName(codeName, messageName string) *HTTPClient {
	c.propertiesCodeName = codeName
	c.propertiesMessageName = messageName
	return c
}

// Limit 限流
//
// Note:
// 1.不调用此方法不启动限流功能
// 2.WithEnabled 设置 false 同样不启动限流功能
// 3.限流配置由项目提前进行初始化
func (c *HTTPClient) Limit(options ...LimitOption) *HTTPClient {
	opts := &limitOptions{}
	opts.applyOpts(options)
	c.limitOpts = opts
	return c
}

// check 参数检查及默认值初始化
func (c *HTTPClient) check() error {
	if c.context == nil {
		return errors.New("the context invalid")
	}

	// 默认重试两次
	if c.retryCount == 0 {
		c.retryCount = 2
	}

	// 禁止重试
	if c.retryCount < 0 {
		c.retryCount = 0
	}

	if c.timeout == 0 {
		c.timeout = 500 * time.Millisecond
	}

	traceID, ok := c.context.Value(varctx.TraceID).(string)
	if !ok {
		traceID = id.RandomID()
	}
	c.client.SetHeader(varheader.TraceID, traceID)

	env, ok := c.context.Value(varctx.Environment).(string)
	if ok {
		c.client.SetHeader(varheader.Environment, env)
	}

	ver, ok := c.context.Value(varctx.Version).(string)
	if ok {
		c.client.SetHeader(varheader.Version, ver)
	}

	if c.userAgent == "" {
		c.userAgent = DefaultUserAgent
	}
	c.client.SetHeader("User-Agent", c.userAgent)

	if len(c.headers) > 0 {
		c.client.SetHeaders(c.headers)
	}

	if c.transport != nil {
		c.client.SetTransport(c.transport)
	}

	c.client.OnBeforeRequest(c.OnBeforeRequest())
	c.client.OnError(c.LimitOnError())

	if c.logger != nil {
		c.client.
			OnError(c.OnError()).
			AddRetryHook(c.OnRetry()).
			OnAfterResponse(c.OnAfterResponse()).
			OnAfterResponse(c.LimitEntryExit())
	}

	// 单机流控
	if c.limitOpts != nil &&
		c.limitOpts.Enabled &&
		c.limitOpts.Type != xlimit.RuleTypeCFR {
		c.client.OnBeforeRequest(c.LimitRequest())
	}

	// 集群流控
	if c.limitOpts != nil &&
		c.limitOpts.Enabled &&
		c.limitOpts.Type == xlimit.RuleTypeCFR &&
		c.limitOpts.RedisLimiter != nil &&
		c.limitOpts.Threshold > 0 &&
		c.limitOpts.StatInterval > 0 {
		c.client.OnBeforeRequest(c.ClusterLimitRequest())
	}

	if c.authOption != nil && c.authOption.appID != "" && c.authOption.appSecret != "" {
		c.client.OnBeforeRequest(c.ApiAuthRequest())
	}

	// Note: 重要日志通过 Hook 实现打印，其它日志不打印
	c.client.SetLogger(NewNopRestyLogger())

	return nil
}

func (c *HTTPClient) Request() *resty.Request {
	err := c.check()
	if err != nil {
		return c.client.R().SetError(err)
	}

	return c.client.
		EnableTrace().
		SetTimeout(c.timeout).
		SetRetryCount(c.retryCount).
		R().
		SetContext(c.context)
}

// TLSClientConfig
func (c *HTTPClient) TLSClientConfig(config *tls.Config) *HTTPClient {
	c.client.SetTLSClientConfig(config)
	return c
}
