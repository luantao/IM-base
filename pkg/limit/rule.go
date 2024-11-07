package limit

import (
	"MyIM/pkg/config"
	"fmt"
	"strings"

	"github.com/alibaba/sentinel-golang/core/circuitbreaker"
	"github.com/alibaba/sentinel-golang/core/flow"
	"github.com/alibaba/sentinel-golang/core/hotspot"
	"github.com/alibaba/sentinel-golang/core/isolation"
	"github.com/alibaba/sentinel-golang/core/system"
)

// 配置规则类型
const (
	RuleTypeFR = "fr" // 流量控制规则 flow rule
	RuleTypeHR = "hr" // 热点参数流控规则 hotspot rule
	RuleTypeIR = "ir" // 并发隔离控制规则 isolation rule
	RuleTypeCR = "cr" // 熔断规则 circuit breaker rule
	RuleTypeSR = "sr" // 系统自适应保护规则 system rule

	RuleTypeCFR = "cfr" // Redis 流控规则 cluster flow rule
)

// Rules 规则集合
type Rules struct {
	flowRules           map[string]*flow.Rule
	hotspotRules        map[string]*hotspot.Rule
	isolationRules      map[string]*isolation.Rule
	circuitBreakerRules map[string]*circuitbreaker.Rule
	systemRules         map[string]*system.Rule
}

// NewRules 全部规则集合
func NewRules() *Rules {
	return &Rules{
		flowRules:           make(map[string]*flow.Rule),
		hotspotRules:        make(map[string]*hotspot.Rule),
		isolationRules:      make(map[string]*isolation.Rule),
		circuitBreakerRules: make(map[string]*circuitbreaker.Rule),
		systemRules:         make(map[string]*system.Rule),
	}
}

// Reset 清空所有规则
//
// 在添加第一条规则前先确保没有任何规则存在
// 如: 规则配置变更后重新加载规则
func (rs *Rules) Reset() {
	rs.flowRules = map[string]*flow.Rule{}
	rs.hotspotRules = map[string]*hotspot.Rule{}
	rs.isolationRules = map[string]*isolation.Rule{}
	rs.circuitBreakerRules = map[string]*circuitbreaker.Rule{}
	rs.systemRules = map[string]*system.Rule{}
}

// AddRules 添加多条规则
// Note:
// 1.覆盖已存在的规则
// 2.若无配置则不进行加载
func (rs *Rules) AddRules(section string) error {
	// 未配置则不进行加载
	if !config.GetIsExist(section) {
		return nil
	}

	buildKey := func(key string) string {
		return strings.Join([]string{section, key}, ".")
	}

	for k, _ := range config.GetStringMap(section) {
		err := rs.AddRule(buildKey(k))
		if err != nil {
			return err
		}
	}
	return nil
}

// AddRule 添加单条规则
//
// Note:
// 1.覆盖已存在的规则
// 2.若无配置则不进行加载
func (rs *Rules) AddRule(section string) error {
	// 未配置则不进行加载
	if !config.GetIsExist(section) {
		return nil
	}

	buildKey := func(key string) string {
		return strings.Join([]string{section, key}, ".")
	}

	ruleType := config.GetString(buildKey("type"))
	if ruleType == "" {
		return fmt.Errorf("limiter: %s, rule type not set", section)
	}

	if ruleType == RuleTypeFR {
		rule, err := NewFRConfigWithACM(section)
		if err != nil {
			return err
		}

		rs.flowRules[rule.Resource] = &rule
		return nil
	}

	if ruleType == RuleTypeHR {
		rule, err := NewHRConfigWithACM(section)
		if err != nil {
			return err
		}

		rs.hotspotRules[rule.Resource] = &rule
		return nil
	}

	if ruleType == RuleTypeIR {
		rule, err := NewIRConfigWithACM(section)
		if err != nil {
			return err
		}

		rs.isolationRules[rule.Resource] = &rule
		return nil
	}

	if ruleType == RuleTypeCR {
		rule, err := NewCRConfigWithACM(section)
		if err != nil {
			return err
		}

		rs.circuitBreakerRules[rule.Resource] = &rule
		return nil
	}

	if ruleType == RuleTypeSR {
		rule, err := NewSRConfigWithACM(section)
		if err != nil {
			return err
		}

		rs.systemRules[rule.MetricType.String()] = &rule
		return nil
	}

	return nil
}

// LoadRules 初始化所有规则
//
// Note: 在添加完所有规则后执行加载所有规则
func (rs *Rules) LoadRules() error {
	// 流量控制
	if len(rs.flowRules) > 0 {
		var rules []*flow.Rule
		for _, rule := range rs.flowRules {
			rules = append(rules, rule)
		}
		if _, err := flow.LoadRules(rules); err != nil {
			return err
		}
	}

	// 热点参数流控
	if len(rs.hotspotRules) > 0 {
		var rules []*hotspot.Rule
		for _, rule := range rs.hotspotRules {
			rules = append(rules, rule)
		}
		if _, err := hotspot.LoadRules(rules); err != nil {
			return err
		}
	}

	// 并发隔离控制
	if len(rs.isolationRules) > 0 {
		var rules []*isolation.Rule
		for _, rule := range rs.isolationRules {
			rules = append(rules, rule)
		}
		if _, err := isolation.LoadRules(rules); err != nil {
			return err
		}
	}

	// 熔断
	if len(rs.circuitBreakerRules) > 0 {
		var rules []*circuitbreaker.Rule
		for _, rule := range rs.circuitBreakerRules {
			rules = append(rules, rule)
		}
		if _, err := circuitbreaker.LoadRules(rules); err != nil {
			return err
		}
	}

	// 系统自适应保护
	if len(rs.systemRules) > 0 {
		var rules []*system.Rule
		for _, rule := range rs.systemRules {
			rules = append(rules, rule)
		}
		if _, err := system.LoadRules(rules); err != nil {
			return err
		}
	}

	return nil
}
