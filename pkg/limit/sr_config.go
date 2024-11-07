package limit

import (
	"fmt"
	"github.com/luantao/IM-base/pkg/config"

	"github.com/alibaba/sentinel-golang/core/system"
)

// SRConfig 系统自适应保护规则 system rule
//
// Doc: https://sentinelguard.io/zh-cn/docs/golang/system-adaptive-protection.html
// 系统保护规则是应用整体维度的，而不是单个调用维度的，并且仅对入口流量生效。
//
// yaml example:
// cpu_load_limiter:
//
//	type: sr # 系统自适应保护
//	enabled: true # 是否启用
//	metric_type: 0 # 指标类型;0:Load;1:AvgRT;2:Concurrency;3:InboundQPS;4:CpuUsage
//	trigger_count: 5 # 阈值
//	strategy: 0 # 策略;-1:无;0:BBR
type SRConfig struct {
	Type         string                  `json:"type"`
	MetricType   system.MetricType       `json:"metricType"`
	TriggerCount float64                 `json:"triggerCount"`
	Strategy     system.AdaptiveStrategy `json:"strategy"`
}

func NewSRConfigWithACM(section string) (rule system.Rule, err error) {
	var conf SRConfig

	err = config.UnmarshalKey(section, &conf)
	if err != nil {
		err = fmt.Errorf("limiter: %s, rule config invalid", section)
		return
	}

	if conf.MetricType == 0 || conf.MetricType > system.CpuUsage {
		err = fmt.Errorf("limiter: %s, rule resource invalid", section)
		return
	}

	rule = system.Rule{
		MetricType:   system.MetricType(conf.MetricType),
		TriggerCount: conf.TriggerCount,
		Strategy:     system.AdaptiveStrategy(conf.Strategy),
	}
	return
}
