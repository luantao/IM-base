package http

import (
	"MyIM/pkg/auth"
	xlimit "MyIM/pkg/limit"
	"MyIM/pkg/mlog"
	"MyIM/pkg/rediss"
	varheader "MyIM/pkg/var/header"
	"errors"
	"fmt"
	"strings"

	sentinel "github.com/alibaba/sentinel-golang/api"
	sentinelbase "github.com/alibaba/sentinel-golang/core/base"
	"github.com/go-resty/resty/v2"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func (hc *HTTPClient) OnBeforeRequest() resty.RequestMiddleware {
	return func(c *resty.Client, r *resty.Request) error {
		urlArr := strings.Split(r.URL, "?")
		hc.uri = urlArr[0]
		return nil
	}
}

func (hc *HTTPClient) OnAfterResponse() resty.ResponseMiddleware {
	return func(c *resty.Client, r *resty.Response) error {
		span := opentracing.SpanFromContext(r.Request.Context())
		if span != nil {
			span.Finish()
		}

		if hc.logger == nil {
			return nil
		}

		if !mlog.Enabled(hc.logger.level, zapcore.InfoLevel) {
			return nil
		}

		respBody := r.String()
		fields := []zapcore.Field{
			zap.String(mlog.NameKey, "http"),
			zap.Float64("latency", r.Time().Seconds()*1000),
			zap.String("request_uri", hc.uri),
			zap.String("host", r.RawResponse.Request.Host),
			zap.String("x_request_id", r.Header().Get(varheader.RequestID)),
			zap.String("request_method", r.Request.Method),
			zap.Float64("timeout", c.GetClient().Timeout.Seconds()*1000),
			zap.Int("status", r.StatusCode()),
			zap.Int("request_attempt", r.Request.Attempt),
		}
		if hc.logRequest {
			fields = append(fields, zap.String("request", r.Request.QueryParam.Encode()))
			fields = append(fields, zap.String("request_formdata", r.Request.FormData.Encode()))
			fields = append(fields, zap.Any("request_body", r.Request.Body))
		}

		if hc.logResponse {
			fields = append(fields, zap.String("response", respBody))
		}

		if hc.addProperties {
			fields = append(fields, zap.Any(mlog.ContentKey, hc.parsePartResponse(respBody)))
		}

		if hc.trace {
			fields = append(fields,
				zap.Float64("dns_Lookup", r.Request.TraceInfo().DNSLookup.Seconds()*1000),
				zap.Float64("conn_time", r.Request.TraceInfo().ConnTime.Seconds()*1000),
				zap.Float64("tcp_conn_time", r.Request.TraceInfo().TCPConnTime.Seconds()*1000),
				zap.Float64("tls_handshake", r.Request.TraceInfo().TLSHandshake.Seconds()*1000),
				zap.Float64("server_time", r.Request.TraceInfo().ServerTime.Seconds()*1000),
				zap.Float64("response_time", r.Request.TraceInfo().ResponseTime.Seconds()*1000),
				zap.Bool("is_conn_reused", r.Request.TraceInfo().IsConnReused),
				zap.Bool("is_conn_was_idle", r.Request.TraceInfo().IsConnWasIdle),
				zap.Float64("conn_idle_time", r.Request.TraceInfo().ConnIdleTime.Seconds()*1000),
				zap.String("remote_addr", r.Request.TraceInfo().RemoteAddr.String()),
				zap.Int("request_attempt", r.Request.TraceInfo().RequestAttempt),
				zap.Float64("total_time", r.Request.TraceInfo().TotalTime.Seconds()*1000),
			)
		}

		hc.logger.WithCTX(r.Request.Context()).Info("http request", fields...)

		//针对非2xx返回，直接报错，防止业务解析非200 response body失败
		if !r.IsSuccess() && hc.checkHttpCode {
			return errors.New("request failed, because the response status code not between 200 and 399")
		}
		return nil
	}
}

func (hc *HTTPClient) OnError() resty.ErrorHook {
	return func(r *resty.Request, err error) {
		span := opentracing.SpanFromContext(r.Context())
		if span != nil {
			span.Finish()
		}

		if hc.logger == nil {
			return
		}

		if !mlog.Enabled(hc.logger.level, zapcore.ErrorLevel) {
			return
		}

		if resErr, ok := err.(*resty.ResponseError); ok {
			latency := resErr.Response.Time().Seconds() * 1000
			if hc.context.Err() != nil {
				latency = float64(resErr.Response.ReceivedAt().UnixNano()-resErr.Response.Request.Time.UnixNano()) / 1e6
			}
			fields := []zapcore.Field{
				zap.String(mlog.NameKey, "http"),
				zap.Error(err),
				zap.Float64("latency", latency),
				zap.String("request_uri", hc.uri),
				zap.String("host", r.RawRequest.Host),
				zap.String("request_method", r.Method),
				zap.String("request", r.QueryParam.Encode()),
				zap.String("request_formdata", r.FormData.Encode()),
				zap.Any("request_body", r.Body),
				zap.Int("request_attempt", r.Attempt),
			}

			if hc.trace {
				fields = append(fields,
					zap.Float64("dns_Lookup", r.TraceInfo().DNSLookup.Seconds()*1000),
					zap.Float64("conn_time", r.TraceInfo().ConnTime.Seconds()*1000),
					zap.Float64("tcp_conn_time", r.TraceInfo().TCPConnTime.Seconds()*1000),
					zap.Float64("tls_handshake", r.TraceInfo().TLSHandshake.Seconds()*1000),
					zap.Float64("server_time", r.TraceInfo().ServerTime.Seconds()*1000),
					zap.Float64("response_time", r.TraceInfo().ResponseTime.Seconds()*1000),
					zap.Bool("is_conn_reused", r.TraceInfo().IsConnReused),
					zap.Bool("is_conn_was_idle", r.TraceInfo().IsConnWasIdle),
					zap.Float64("conn_idle_time", r.TraceInfo().ConnIdleTime.Seconds()*1000),
					// zap.Int("request_attempt", r.TraceInfo().RequestAttempt),
					// zap.Float64("total_time", r.TraceInfo().TotalTime.Seconds()*1000),
				)
			}

			hc.logger.WithCTX(r.Context()).Error("http request", fields...)
			return
		}

		hc.logger.WithCTX(r.Context()).Error("http request",
			zap.String(mlog.NameKey, "http"),
			zap.String("request_uri", hc.uri),
			zap.Error(err),
		)
	}
}

func (hc *HTTPClient) OnRetry() resty.OnRetryFunc {
	return func(r *resty.Response, err error) {
		if hc.logger == nil {
			return
		}

		if !mlog.Enabled(hc.logger.level, zapcore.WarnLevel) {
			return
		}

		fields := []zapcore.Field{
			zap.String(mlog.NameKey, "http"),
			zap.Error(err),
			zap.Float64("latency", r.Time().Seconds()*1000),
			zap.String("request_uri", hc.uri),
			zap.String("host", r.Request.RawRequest.Host),
			zap.String("x_request_id", r.Header().Get(varheader.RequestID)),
			zap.String("request_method", r.Request.Method),
			zap.String("request", r.Request.QueryParam.Encode()),
			zap.String("request_formdata", r.Request.FormData.Encode()),
			zap.Any("request_body", r.Request.Body),
			zap.Int("request_attempt", r.Request.Attempt),
		}

		if hc.trace {
			fields = append(fields,
				zap.Float64("dns_Lookup", r.Request.TraceInfo().DNSLookup.Seconds()*1000),
				zap.Float64("conn_time", r.Request.TraceInfo().ConnTime.Seconds()*1000),
				zap.Float64("tcp_conn_time", r.Request.TraceInfo().TCPConnTime.Seconds()*1000),
				zap.Float64("tls_handshake", r.Request.TraceInfo().TLSHandshake.Seconds()*1000),
				zap.Float64("server_time", r.Request.TraceInfo().ServerTime.Seconds()*1000),
				zap.Float64("response_time", r.Request.TraceInfo().ResponseTime.Seconds()*1000),
				zap.Bool("is_conn_reused", r.Request.TraceInfo().IsConnReused),
				zap.Bool("is_conn_was_idle", r.Request.TraceInfo().IsConnWasIdle),
				zap.Float64("conn_idle_time", r.Request.TraceInfo().ConnIdleTime.Seconds()*1000),
			)
		}

		hc.logger.WithCTX(r.Request.Context()).Warn("http request", fields...)
	}
}

// LimitOnError 请求失败后设置限流错误
func (hc *HTTPClient) LimitOnError() resty.ErrorHook {
	return func(r *resty.Request, err error) {
		if hc.limitEntry == nil {
			return
		}
		if err != nil {
			hc.limitEntry.SetError(err)
			hc.limitEntry.Exit()
		}
		return
	}
}

// LimitOnAfter 请求完成后设置限流错误
func (hc *HTTPClient) LimitEntryExit() resty.ResponseMiddleware {
	return func(c *resty.Client, r *resty.Response) error {
		if hc.limitEntry != nil {
			hc.limitEntry.Exit()
		}
		return nil
	}
}

// LimitRequest sentinel 单机请求限流
func (hc *HTTPClient) LimitRequest() resty.RequestMiddleware {
	return func(c *resty.Client, r *resty.Request) error {
		if hc.limitOpts.ResourceName == "" {
			hc.limitOpts.ResourceName = r.Method + ":" + hc.uri
		}

		// 限流埋点
		entry, err := sentinel.Entry(
			hc.limitOpts.ResourceName,
			sentinel.WithTrafficType(sentinelbase.Outbound),
			sentinel.WithResourceType(sentinelbase.ResTypeWeb),
			sentinel.WithArgs(hc.limitOpts.HotKeys...),
		)
		hc.limitEntry = entry
		// 未触发限流
		if err == nil {
			return nil
		}
		return err
	}
}

// ClusterLimitRequest redis 集群限流
func (hc *HTTPClient) ClusterLimitRequest() resty.RequestMiddleware {
	return func(c *resty.Client, r *resty.Request) error {
		key := hc.limitOpts.ResourceName
		if key == "" {
			key = r.Method + ":" + hc.uri
		}

		if hc.limitOpts.KeyPrefix != "" {
			key = hc.limitOpts.KeyPrefix + ":" + key
		}

		for _, hotkey := range hc.limitOpts.HotKeys {
			key = key + ":" + fmt.Sprintf("%v", hotkey)
		}

		// 限流埋点
		limit := rediss.Limit{
			Rate:   hc.limitOpts.Threshold,
			Burst:  hc.limitOpts.Threshold,
			Period: hc.limitOpts.StatInterval,
		}

		result, err := hc.limitOpts.RedisLimiter.AllowN(hc.context, key, limit, 1)
		hc.logger.WithCTX(hc.context).Debug("redis cluster limit result",
			zap.Any("result", result),
		)

		if err != nil {
			return err
		}

		// 触发限流
		if result.Allowed == 0 || result.RetryAfter > -1 {
			return &xlimit.ClusterLimitError{
				ErrMsg: "trigger cluster current limit",
				Result: result,
			}
		}

		// 未触发限流
		return nil
	}
}

// ApiAuthRequest api 鉴权
func (hc *HTTPClient) ApiAuthRequest() resty.RequestMiddleware {
	return func(c *resty.Client, r *resty.Request) error {
		r.SetHeader(varheader.Authorization,
			auth.GenAuthorizationInfo(hc.authOption.appID, hc.authOption.appSecret))
		return nil
	}
}
