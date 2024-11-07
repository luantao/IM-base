package limit

// CFRConfig 集群限流规则 cluster flow rule
//
// Doc: https://github.com/go-redis/redis_rate
//
// 使用 Redis 进行集群限流方案
// 以后 sentinel 发布集群限流功能后再单独接入
// CFR 规则配置与 sentinel 无关，在此声明，主要是为了集中展示流控支持的所有规则。
//
// yaml example:
// order_center:
//   order_detail:
//     url: 'http://xx.xxx.com/order/v1/agent/order/show'
//     timeout: 100 # 单位毫秒
//     limiter:
//       type: cfr # Redis 集群限流
//       enabled: true # 是否启用
//       resource: 'create-order' # 唯一标识;redis key 的一部分,尽量简短
//       key_prefix: '' # 默认无;通常为项目名字减少 key 冲突
//       threshold: 1000 # 单位时间内(stat_interval_in_ms)请求数量;0值不限流
//       stat_interval_in_ms: 1000 # 统计周期;单位:毫秒;0值不限流
