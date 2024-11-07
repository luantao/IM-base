package http

import (
	"github.com/luantao/IM-base/pkg/config"
	"github.com/luantao/IM-base/pkg/rediss"
	"strings"
	"time"
)

// limitOptions 可选参数选项
type limitOptions struct {
	// 以下是 sentinel 和 redis 集群限流共有配置
	Type         string        `json:"type"`          // 限流规则类型;参考 limit.RuleType
	Enabled      bool          `json:"enabled"`       // 是否开启
	ResourceName string        `json:"resource_name"` // 资源唯一标识
	HotKeys      []interface{} `json:"hot_keys"`      // 热点参数统计字段

	// 以下是 Redis 集群限流独有配置(type 为 cfr)
	RedisLimiter *rediss.Limiter // redis 限流器
	KeyPrefix    string          // redis key 前缀
	Threshold    int             // 单位时间内(stat_interval_in_ms)请求数量
	StatInterval time.Duration   // 统计周期;单位:毫秒
}

// LimitOption 可选参数类型
type LimitOption func(*limitOptions)

// applyOpts 应用可选参数
func (op *limitOptions) applyOpts(opts []LimitOption) {
	for _, opt := range opts {
		opt(op)
	}
}

type limit struct{}

// Limit 限流参数集合
func Limit() limit { return limit{} }

// Config 从 Config 中读取相关配置
// 其配置信息见 limit 模块各种 rule config
//
// 可省略调用 Enabled 和 Name
func (limit) Config(section string) LimitOption {
	buildKey := func(key string) string {
		return strings.Join([]string{section, key}, ".")
	}

	return func(opts *limitOptions) {
		opts.Type = config.GetString(buildKey("type"))
		opts.Enabled = config.GetBool(buildKey("enabled"))
		opts.ResourceName = config.GetString(buildKey("resource"))
		opts.KeyPrefix = config.GetString(buildKey("key_prefix"))
		opts.Threshold = config.GetInt(buildKey("threshold"))
		opts.StatInterval = config.GetDuration(buildKey("stat_interval_in_ms")) * time.Millisecond
	}
}

// Enabled 是否开启
//
// 尽量通过 acm 配置
// Deprecated: 使用 Config 方法代替.
func (limit) Enabled(enabled bool) LimitOption {
	return func(opts *limitOptions) { opts.Enabled = enabled }
}

// Name 限流器名字(资源统计唯一标识)
//
// Deprecated: 使用 Config 方法代替.
//
// 1.sentinel 限流时:
// @param resourceName string 可为空可定义
// 必须和 acm 规则中一致
// 为空时默认格式(Method:URI) e.g. GET:https://www.xx.com/order/show
//
// 2.redis 限流时:
// 配置规则同上，唯一不同的是: resourceName 将作为 redis key 的一部分
func (limit) Name(resourceName string) LimitOption {
	return func(opts *limitOptions) { opts.ResourceName = resourceName }
}

// HotKeys 热 key 埋点
//
// 1.sentinel 限流时对应 sentinel WithArgs
//
// 2.redis 限流时将联合作为 redis key 的一部分
// 如：rate:oc:dic:open_city:hotkey1(渠道ID):hotkey2(城市ID)
func (limit) HotKeys(keys ...interface{}) LimitOption {
	return func(opts *limitOptions) { opts.HotKeys = append(opts.HotKeys, keys...) }
}

// Cluster Redis 集群限流器
//
// 若未配置 CFR 限流器，即时使用也无效果，反之亦然
func (limit) Cluster(redisLimiter *rediss.Limiter) LimitOption {
	return func(opts *limitOptions) { opts.RedisLimiter = redisLimiter }
}
