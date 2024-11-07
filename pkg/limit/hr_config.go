package limit

import (
	"MyIM/pkg/config"
	"fmt"

	"github.com/alibaba/sentinel-golang/core/hotspot"
)

// HRConfig 热点参数流控规则 hotspot rule
//
// Doc: https://sentinelguard.io/zh-cn/docs/golang/hotspot-param-flow-control.html
//
// yaml example:
// limiter:
//
//	type: hr # 热点参数流控
//	enabled: true # 是否启用
//	resource: 'create-order' # 唯一标识
//	metric_type: # 指标类型;0:并发;1:qps
//	control_behavior: 0 # 控制策略;0:Reject(拒绝);1:Throttling(匀速)
//	param_index: 0 # 热点参数索引
//	param_key: '' # 热点参数字段
//	threshold: 1000 # 限流阈值（针对每个热点参数）
//	# max_queueing_time_ms: 0 # 最大排队等待时长
//	# burst_count: 0 # 静默值(仅在快速失败模式 + QPS 下生效)
//	# duration_in_sec: 0 # 统计结构填充新的 token 的时间间隔 (仅在请求数(QPS)流控模式下生效)
//	# params_max_capacity: 20000 # 统计结构的容量最大值（Top N）
type HRConfig struct {
	Type              string                  `json:"type"`
	ID                string                  `json:"id,omitempty" mapstructure:"id"`
	Resource          string                  `json:"resource" mapstructure:"resource"`
	MetricType        hotspot.MetricType      `json:"metric_type" mapstructure:"metric_type"`
	ControlBehavior   hotspot.ControlBehavior `json:"control_behavior" mapstructure:"control_behavior"`
	ParamIndex        int                     `json:"param_index" mapstructure:"param_index"`
	ParamKey          string                  `json:"param_key" mapstructure:"param_key"`
	Threshold         int64                   `json:"threshold" mapstructure:"threshold"`
	MaxQueueingTimeMs int64                   `json:"max_queueing_time_ms" mapstructure:"max_queueing_time_ms"`
	BurstCount        int64                   `json:"burst_count" mapstructure:"burst_count"`
	DurationInSec     int64                   `json:"duration_in_sec" mapstructure:"duration_in_sec"`
	ParamsMaxCapacity int64                   `json:"params_max_capacity" mapstructure:"params_max_capacity"`
}

func NewHRConfigWithACM(section string) (rule hotspot.Rule, err error) {
	var conf HRConfig

	err = config.UnmarshalKey(section, &conf)
	if err != nil {
		err = fmt.Errorf("limiter: %s, rule config invalid", section)
		return
	}

	if conf.Resource == "" {
		err = fmt.Errorf("limiter: %s, rule resource invalid", section)
		return
	}

	rule = hotspot.Rule{
		ID:                conf.ID,
		Resource:          conf.Resource,
		MetricType:        hotspot.MetricType(conf.MetricType),
		ControlBehavior:   hotspot.ControlBehavior(conf.ControlBehavior),
		ParamIndex:        conf.ParamIndex,
		ParamKey:          conf.ParamKey,
		Threshold:         conf.Threshold,
		MaxQueueingTimeMs: conf.MaxQueueingTimeMs,
		BurstCount:        conf.BurstCount,
		DurationInSec:     conf.DurationInSec,
		ParamsMaxCapacity: conf.ParamsMaxCapacity,
	}
	return
}
