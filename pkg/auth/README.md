# API统一鉴权

> [前期设计文档](https://yangguangchuxing.feishu.cn/wiki/wikcnORIiebiEg7jbJluN1vm9rh)

## 最佳实践

* 一个 `app_id` 可以应用于一个服务下的所有接口、一组接口，或者加在特定的接口上面
* 一个服务可以有多个 `app_id`

## 注意事项

不要将服务A申请的 `app_id` 给到服务B，一个 `app_id` 通常通过鉴权后会对应一套限流规则，下游服务通过相应的监控会注意到这种问题并推动调用方调整。


## 限流中间件配置样例

```yaml
service_rule:
    {{uri}}:
      {{app_id01}}:
      rule: selte.rule_key_1
      disabled: true
      default: # 缺省
        rule: ""
        disabled: false
      default:
          {{app_id01}}:
            rule: selte.rule_key_1
            disabled: true
            default: # 缺省
              rule: ""
              disabled: false


selte:
  rule_key_1:
    type: fr # 流量控制
    enabled: true # 是否启用
    resource: 'create-order' # 唯一标识
    token_calculate_strategy: 0 # Token计算策略;0:Direct;1:WarmUp;2:MemoryAdaptive
    control_behavior: 0 # 控制策略;0:Reject(拒绝);1:Throttling(匀速)
    threshold: 1000 # 单位时间内(stat_interval_in_ms)请求数量
    relation_strategy: # 调用关系限流策略;0:CurrentResource;1:AssociatedResource
    ref_resource: '' # 关联资源
    max_queueing_time_ms: 100 # 匀速排队的最大等待时间
    warm_up_period_sec: 300 # 预热的时间长度
    warm_up_cold_factor: 3 # 预热的因子，默认是3
    stat_interval_in_ms: 1000 # 统计周期
    # low_mem_usage_threshold:
    # high_mem_usage_threshold:
    # mem_low_water_mark_bytes:
    # mem_high_water_mark_bytes:
  rule_key_2:
    type: fr # 流量控制
    enabled: true # 是否启用
    resource: 'create-order' # 唯一标识
    token_calculate_strategy: 0 # Token计算策略;0:Direct;1:WarmUp;2:MemoryAdaptive
    control_behavior: 0 # 控制策略;0:Reject(拒绝);1:Throttling(匀速)
    threshold: 1000 # 单位时间内(stat_interval_in_ms)请求数量
    relation_strategy: # 调用关系限流策略;0:CurrentResource;1:AssociatedResource
    ref_resource: '' # 关联资源
    max_queueing_time_ms: 100 # 匀速排队的最大等待时间
    warm_up_period_sec: 300 # 预热的时间长度
    warm_up_cold_factor: 3 # 预热的因子，默认是3
    stat_interval_in_ms: 1000 # 统计周期
    # low_mem_usage_threshold:
    # high_mem_usage_threshold:
    # mem_low_water_mark_bytes:
    # mem_high_water_mark_bytes:    
```

