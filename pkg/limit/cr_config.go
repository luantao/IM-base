package limit

import (
	"fmt"
	"github.com/luantao/IM-base/pkg/config"

	"github.com/alibaba/sentinel-golang/core/circuitbreaker"
)

// CRConfig 熔断规则 circuit breaker rule
//
// Doc: https://sentinelguard.io/zh-cn/docs/golang/circuit-breaking.html
//
// yaml example:
// order_center:
//
//	order_detail:
//	  url: 'http://xx.xxx.com/order/v1/agent/order/show'
//	  timeout: 100 # 单位毫秒
//	  limiter:
//	    type: cr # 熔断器
//	    enabled: true # 是否启用
//	    resource: 'oc:order/v1/agent/order/show' # 唯一标识
//	    strategy: 1 # 策略;0:慢请求比例;1:错误比例;2:错误数
//	    threshold: 0.3 # 阈值;慢请求比例阈值或错误比例阈值
//	    retry_timeout_ms: 5000 # 熔断触发后持续的时间
//	    min_request_amount: 100 # 触发熔断的最小请求数
//	    stat_interval_ms: 1000 # 统计的时间窗口长度
//	    max_allowed_rt_ms: 1000 # 慢请求时间;请求时间大于此值属于慢请求
//	    # stat_sliding_window_bucket_count: 1
type CRConfig struct {
	Type                         string                  `json:"type"`
	ID                           string                  `json:"id,omitempty" mapstructure:"id"`
	Resource                     string                  `json:"resource" mapstructure:"resource"`
	Strategy                     circuitbreaker.Strategy `json:"strategy" mapstructure:"strategy"`
	RetryTimeoutMs               uint32                  `json:"retry_timeout_ms" mapstructure:"retry_timeout_ms"`
	MinRequestAmount             uint64                  `json:"min_request_amount" mapstructure:"min_request_amount"`
	StatIntervalMs               uint32                  `json:"stat_interval_ms" mapstructure:"stat_interval_ms"`
	StatSlidingWindowBucketCount uint32                  `json:"stat_sliding_window_bucket_count" mapstructure:"stat_sliding_window_bucket_count"`
	MaxAllowedRtMs               uint64                  `json:"max_allowed_rt_ms" mapstructure:"max_allowed_rt_ms"`
	Threshold                    float64                 `json:"threshold" mapstructure:"threshold"`
}

func NewCRConfigWithACM(section string) (rule circuitbreaker.Rule, err error) {
	var conf CRConfig

	err = config.UnmarshalKey(section, &conf)
	if err != nil {
		err = fmt.Errorf("limiter: %s, rule config invalid", section)
		return
	}

	if conf.Resource == "" {
		err = fmt.Errorf("limiter: %s, rule resource invalid", section)
		return
	}

	rule = circuitbreaker.Rule{
		Id:                           conf.ID,
		Resource:                     conf.Resource,
		Strategy:                     circuitbreaker.Strategy(conf.Strategy),
		RetryTimeoutMs:               conf.RetryTimeoutMs,
		MinRequestAmount:             conf.MinRequestAmount,
		StatIntervalMs:               conf.StatIntervalMs,
		StatSlidingWindowBucketCount: conf.StatSlidingWindowBucketCount,
		MaxAllowedRtMs:               conf.MaxAllowedRtMs,
		Threshold:                    conf.Threshold,
	}
	return
}
