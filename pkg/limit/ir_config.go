package limit

import (
	"fmt"
	"github.com/luantao/IM-base/pkg/config"

	"github.com/alibaba/sentinel-golang/core/isolation"
)

// IRConfig 并发隔离控制规则 isolation rule
//
// Doc: https://sentinelguard.io/zh-cn/docs/golang/concurrency-limiting-isolation.html
//
// yaml example:
// limiter:
//
//	type: ir # 并发隔离控制
//	enabled: true # 是否启用
//	resource: 'create-order' # 唯一标识
//	metric_type: 0 # 指标类型;0:并发
//	threshold: 100 # 限流阈值
type IRConfig struct {
	Type       string               `json:"type"`
	ID         string               `json:"id,omitempty" mapstructure:"id"`
	Resource   string               `json:"resource" mapstructure:"resource"`
	MetricType isolation.MetricType `json:"metric_type" mapstructure:"metric_type"`
	Threshold  uint32               `json:"threshold" mapstructure:"threshold"`
}

func NewIRConfigWithACM(section string) (rule isolation.Rule, err error) {
	var conf IRConfig

	err = config.UnmarshalKey(section, &conf)
	if err != nil {
		err = fmt.Errorf("limiter: %s, rule config invalid", section)
		return
	}

	if conf.Resource == "" {
		err = fmt.Errorf("limiter: %s, rule resource invalid", section)
		return
	}

	rule = isolation.Rule{
		ID:         conf.ID,
		Resource:   conf.Resource,
		MetricType: isolation.MetricType(conf.MetricType),
		Threshold:  conf.Threshold,
	}
	return
}
