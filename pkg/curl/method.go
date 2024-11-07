package curl

import (
	"MyIM/pkg/config"
	"MyIM/pkg/http"
	"MyIM/pkg/mlog"
	"context"
	"encoding/json"
	"go.uber.org/zap"
	"strings"
	"time"
)

// LoadConfig 加载配置
func (this *MethodData) LoadConfig(ctx context.Context, conf string) (err error) {
	this.Url = config.GetString(strings.Join([]string{conf, "url"}, "."))
	this.Timeout = config.GetInt64(strings.Join([]string{conf, "timeout"}, "."))
	this.RetryCount = config.GetInt(strings.Join([]string{conf, "retry"}, "."))
	this.Limit = http.Limit().Config(strings.Join([]string{conf, "limiter"}, "."))
	return
}

func (this *MethodData) PostJson(ctx context.Context, param interface{}, response interface{}) (err error) {
	logger := mlog.Logger().WithCTX(ctx)
	if this.Config != "" {
		err = this.LoadConfig(ctx, this.Config)
		if err != nil {
			logger.Error("http.loadConfig.err", zap.Error(err))
			return
		}
	}

	err = this.validation()
	if err != nil {
		logger.Error("http.post.err", zap.Error(err))
		return
	}

	if _, ok := this.Header["Content-Type"]; !ok {
		this.Header["Content-Type"] = "application/json"
	}
	resp, err := http.Default().Client(ctx).
		Headers(this.Header).
		Timeout(time.Duration(this.Timeout) * time.Millisecond).
		RetryCount(this.RetryCount).
		Request().
		SetBody(param).
		Post(this.Url)

	if err != nil {
		logger.Error(
			"http.request.failed",
			zap.Error(err),
			zap.String("request_uri", this.Url),
			zap.Any("req", param),
		)
		return
	}

	err = json.Unmarshal(resp.Body(), response)
	if err != nil {
		mlog.Logger().WithCTX(ctx).Error(
			"http.jsonUnmarshal.failed",
			zap.Error(err),
			zap.String("request_uri", this.Url),
			zap.Any("req", param),
			zap.Any("resp", resp),
		)
		return
	}

	return
}

func (this *MethodData) PostFormData(ctx context.Context, param map[string]string, response interface{}) (err error) {
	logger := mlog.Logger().WithCTX(ctx)

	if this.Config != "" {
		err = this.LoadConfig(ctx, this.Config)
		if err != nil {
			logger.Error("http.loadConfig.err", zap.Error(err))
			return
		}
	}
	err = this.validation()
	if err != nil {
		logger.Error("http.post.err", zap.Error(err))
		return
	}

	if _, ok := this.Header["Content-Type"]; !ok {
		this.Header["Content-Type"] = "multipart/form-data"
	}

	resp, err := http.Default().Client(ctx).
		Headers(this.Header).
		Timeout(time.Duration(this.Timeout) * time.Millisecond).
		RetryCount(this.RetryCount).
		Request().
		SetFormData(param).
		Post(this.Url)

	if err != nil {
		logger.Error(
			"http.request.failed",
			zap.Error(err),
			zap.String("request_uri", this.Url),
			zap.Any("req", param),
		)
		return
	}

	err = json.Unmarshal(resp.Body(), response)
	if err != nil {
		mlog.Logger().WithCTX(ctx).Error(
			"http.jsonUnmarshal.failed",
			zap.Error(err),
			zap.String("request_uri", this.Url),
			zap.Any("req", param),
			zap.Any("resp", resp),
		)
		return
	}

	return
}

func (this *MethodData) Get(ctx context.Context, param map[string]string, response interface{}) (err error) {
	logger := mlog.Logger().WithCTX(ctx)

	if this.Config != "" {
		err = this.LoadConfig(ctx, this.Config)
		if err != nil {
			logger.Error("http.loadConfig.err", zap.Error(err))
			return
		}
	}

	err = this.validation()
	if err != nil {
		logger.Error("http.get.err", zap.Error(err))
		return
	}

	resp, err := http.Default().Client(ctx).
		Headers(this.Header).
		Timeout(time.Duration(this.Timeout) * time.Millisecond).
		RetryCount(this.RetryCount).
		Request().
		SetQueryParams(param).
		Get(this.Url)
	if err != nil {
		logger.Error(
			"http.request.failed",
			zap.Error(err),
			zap.String("request_uri", this.Url),
			zap.Any("req", param),
		)
		return
	}
	err = json.Unmarshal(resp.Body(), response)
	if err != nil {
		mlog.Logger().WithCTX(ctx).Error(
			"http.jsonUnmarshal.failed",
			zap.Error(err),
			zap.String("request_uri", this.Url),
			zap.Any("req", param),
			zap.Any("resp", resp.Body()),
		)
		return
	}
	return
}

func (this *MethodData) Post(ctx context.Context, param interface{}, response interface{}) (err error) {
	logger := mlog.Logger().WithCTX(ctx)

	if this.Config != "" {
		err = this.LoadConfig(ctx, this.Config)
		if err != nil {
			logger.Error("http.loadConfig.err", zap.Error(err))
			return
		}
	}

	err = this.validation()
	if err != nil {
		logger.Error("http.post.err", zap.Error(err))
		return
	}

	resp, err := http.Default().Client(ctx).
		Headers(this.Header).
		Timeout(time.Duration(this.Timeout) * time.Millisecond).
		RetryCount(this.RetryCount).
		Request().
		SetBody(param).
		Post(this.Url)

	if err != nil {
		logger.Error(
			"http.request.failed",
			zap.Error(err),
			zap.String("request_uri", this.Url),
			zap.Any("req", param),
		)
		return
	}

	err = json.Unmarshal(resp.Body(), response)
	if err != nil {
		mlog.Logger().WithCTX(ctx).Error(
			"http.jsonUnmarshal.failed",
			zap.Error(err),
			zap.String("request_uri", this.Url),
			zap.Any("req", param),
			zap.Any("resp", resp),
		)
		return
	}

	return
}