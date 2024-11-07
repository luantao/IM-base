package rocket

import (
	"strings"
	"time"

	acm "MyIM/pkg/config"

	"MyIM/pkg/rocketmq-client-go/consumer"
	"MyIM/pkg/rocketmq-client-go/primitive"
)

// ConsumerOptions 根据配置信息构造消费者参数
func ConsumerOptions(config ConsumerConfig) (ops []consumer.Option) {
	if len(config.Namesrvs) == 0 {
		panic("name servers not set")
	}

	if config.TopicName == "" {
		panic("topic name not set")
	}

	if config.GroupName == "" {
		panic("group name not set")
	}

	if config.InstanceName == "" {
		panic("instance name not set")
	}

	credentials := primitive.Credentials{
		AccessKey:     config.Auth.AK,
		SecretKey:     config.Auth.SK,
		SecurityToken: config.Auth.Token,
	}

	ops = []consumer.Option{
		consumer.WithCredentials(credentials),
		consumer.WithGroupName(config.GroupName),
		consumer.WithConsumerOrder(config.ConsumerOrderly),
		consumer.WithAutoCommit(config.AutoCommit),
		consumer.WithVIPChannel(config.VipChannelEnabled),
		consumer.WithConsumerModel(consumer.MessageModel(config.ConsumerModel)),
	}

	if err := primitive.NamesrvAddr(config.Namesrvs).Check(); err != nil {
		ops = append(ops, consumer.WithNameServerDomain(config.Namesrvs[0]))
	} else {
		ops = append(ops, consumer.WithNameServer(config.Namesrvs))
	}

	if config.InstanceName != "" {
		ops = append(ops, consumer.WithInstance(config.InstanceName))
	} else {
		ops = append(ops, consumer.WithInstance(strings.Join(config.Namesrvs, ";")))
	}

	if config.Namespace != "" {
		ops = append(ops, consumer.WithNamespace(config.Namespace))
	}

	consumeFromWhere := consumer.ConsumeFromWhere(config.ConsumerFromWhere)
	if consumeFromWhere >= consumer.ConsumeFromLastOffset && consumeFromWhere <= consumer.ConsumeFromTimestamp {
		ops = append(ops, consumer.WithConsumeFromWhere(consumeFromWhere))
	}

	if config.MessageBatchMaxSize > 0 {
		ops = append(ops, consumer.WithConsumeMessageBatchMaxSize(config.MessageBatchMaxSize))
	}

	if config.PullBatchSize > 0 {
		ops = append(ops, consumer.WithPullBatchSize(config.PullBatchSize))
	}

	if config.RebalanceLockInterval > 0 {
		ops = append(ops, consumer.WithRebalanceLockInterval(config.RebalanceLockInterval*time.Millisecond))
	}

	if config.SuspendCurrentQueueTime > 0 {
		ops = append(ops, consumer.WithSuspendCurrentQueueTimeMillis(config.SuspendCurrentQueueTime*time.Millisecond))
	}

	if config.PullInterval > 0 {
		ops = append(ops, consumer.WithPullInterval(config.PullInterval*time.Millisecond))
	}

	if config.MaxReconsumeTimes > 0 {
		ops = append(ops, consumer.WithMaxReconsumeTimes(config.MaxReconsumeTimes))
	}

	return
}

// ConsumerOptionsWithACM 根据 ACM 配置信息构造消费者参数
func ConsumerOptionsWithACM(section string) []consumer.Option {
	buildKey := func(key string) string {
		return strings.Join([]string{section, key}, ".")
	}

	consumerModel := 1
	if acm.GetIsExist(buildKey("consumer_model")) {
		consumerModel = acm.GetInt(buildKey("consumer_model"))
	}

	config := ConsumerConfig{
		Namesrvs: acm.GetStringSlice(buildKey("namesrvs")),
		Auth: Auth{
			AK:    acm.GetString(buildKey("auth.ak")),
			SK:    acm.GetString(buildKey("auth.sk")),
			Token: acm.GetString(buildKey("auth.token")),
		},
		TopicName:               acm.GetString(buildKey("topic_name")),
		GroupName:               acm.GetString(buildKey("group_name")),
		Namespace:               acm.GetString(buildKey("namespace")),
		InstanceName:            acm.GetString(buildKey("instance_name")),
		Tags:                    acm.GetString(buildKey("tags")),
		ConsumerModel:           consumerModel,
		ConsumerFromWhere:       acm.GetInt(buildKey("consumer_from_where")),
		ConsumerOrderly:         acm.GetBool(buildKey("consumer_orderly")),
		MaxReconsumeTimes:       acm.GetInt32(buildKey("max_reconsume_times")),
		PullBatchSize:           acm.GetInt32(buildKey("pull_batch_size")),
		MessageBatchMaxSize:     acm.GetInt(buildKey("message_batch_max_size")),
		AutoCommit:              acm.GetBool(buildKey("auto_commit")),
		RebalanceLockInterval:   acm.GetDuration(buildKey("rebalance_lock_interval")),
		SuspendCurrentQueueTime: acm.GetDuration(buildKey("suspend_current_queue_time")),
		PullInterval:            acm.GetDuration(buildKey("pull_interval")),
		VipChannelEnabled:       acm.GetBool(buildKey("vip_channel_enabled")),
	}
	return ConsumerOptions(config)
}
