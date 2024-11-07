package limit

import (
	"MyIM/pkg/config"
	"fmt"

	"github.com/alibaba/sentinel-golang/core/flow"
)

// FRConfig 流量控制规则 flow rule
//
// Doc: https://sentinelguard.io/zh-cn/docs/golang/flow-control.html
//
// yaml example:
// limiter:
//
//	type: fr # 流量控制
//	enabled: true # 是否启用
//	resource: 'create-order' # 唯一标识
//	token_calculate_strategy: 0 # Token计算策略;0:Direct;1:WarmUp;2:MemoryAdaptive
//	control_behavior: 0 # 控制策略;0:Reject(拒绝);1:Throttling(匀速)
//	threshold: 1000 # 单位时间内(stat_interval_in_ms)请求数量
//	relation_strategy: # 调用关系限流策略;0:CurrentResource;1:AssociatedResource
//	ref_resource: '' # 关联资源
//	max_queueing_time_ms: 100 # 匀速排队的最大等待时间
//	warm_up_period_sec: 300 # 预热的时间长度
//	warm_up_cold_factor: 3 # 预热的因子，默认是3
//	stat_interval_in_ms: 1000 # 统计周期
//	# low_mem_usage_threshold:
//	# high_mem_usage_threshold:
//	# mem_low_water_mark_bytes:
//	# mem_high_water_mark_bytes:
type FRConfig struct {
	Type                   string                      `json:"type"`
	ID                     string                      `json:"id,omitempty" mapstructure:"id"`
	Resource               string                      `json:"resource" mapstructure:"resource"`
	TokenCalculateStrategy flow.TokenCalculateStrategy `json:"token_calculate_strategy" mapstructure:"token_calculate_strategy"`
	ControlBehavior        flow.ControlBehavior        `json:"control_behavior" mapstructure:"control_behavior"`
	Threshold              float64                     `json:"threshold" mapstructure:"threshold"`
	RelationStrategy       flow.RelationStrategy       `json:"relation_strategy" mapstructure:"relation_strategy"`
	RefResource            string                      `json:"ref_resource" mapstructure:"ref_resource"`
	MaxQueueingTimeMs      uint32                      `json:"max_queueing_time_ms" mapstructure:"max_queueing_time_ms"`
	WarmUpPeriodSec        uint32                      `json:"warm_up_period_sec" mapstructure:"warm_up_period_sec"`
	WarmUpColdFactor       uint32                      `json:"warm_up_cold_factor" mapstructure:"warm_up_cold_factor"`
	StatIntervalInMs       uint32                      `json:"stat_interval_in_ms" mapstructure:"stat_interval_in_ms"`
	LowMemUsageThreshold   int64                       `json:"low_mem_usage_threshold" mapstructure:"low_mem_usage_threshold"`
	HighMemUsageThreshold  int64                       `json:"high_mem_usage_threshold" mapstructure:"high_mem_usage_threshold"`
	MemLowWaterMarkBytes   int64                       `json:"mem_low_water_mark_bytes" mapstructure:"mem_low_water_mark_bytes"`
	MemHighWaterMarkBytes  int64                       `json:"mem_high_water_mark_bytes" mapstructure:"mem_high_water_mark_bytes"`
}

func NewFRConfigWithACM(section string) (rule flow.Rule, err error) {
	var conf FRConfig

	err = config.UnmarshalKey(section, &conf)
	if err != nil {
		err = fmt.Errorf("limiter: %s, rule config invalid", section)
		return
	}

	if conf.Resource == "" {
		err = fmt.Errorf("limiter: %s, rule resource invalid", section)
		return
	}

	rule = flow.Rule{
		ID:                     conf.ID,
		Resource:               conf.Resource,
		TokenCalculateStrategy: flow.TokenCalculateStrategy(conf.TokenCalculateStrategy),
		ControlBehavior:        flow.ControlBehavior(conf.ControlBehavior),
		Threshold:              conf.Threshold,
		RelationStrategy:       flow.RelationStrategy(conf.RelationStrategy),
		RefResource:            conf.RefResource,
		MaxQueueingTimeMs:      conf.MaxQueueingTimeMs,
		WarmUpPeriodSec:        conf.WarmUpPeriodSec,
		WarmUpColdFactor:       conf.WarmUpColdFactor,
		StatIntervalInMs:       conf.StatIntervalInMs,
		LowMemUsageThreshold:   conf.LowMemUsageThreshold,
		HighMemUsageThreshold:  conf.HighMemUsageThreshold,
		MemLowWaterMarkBytes:   conf.MemLowWaterMarkBytes,
		MemHighWaterMarkBytes:  conf.MemHighWaterMarkBytes,
	}
	return
}
