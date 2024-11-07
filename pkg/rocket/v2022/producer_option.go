package rocket

import (
	"strings"
	"time"

	acm "MyIM/pkg/config"

	"MyIM/pkg/rocketmq-client-go/primitive"
	"MyIM/pkg/rocketmq-client-go/producer"
)

// ProducerOptions 根据配置信息构造生产者参数
func ProducerOptions(config ProducerConfig) (ops []producer.Option) {
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

	ops = []producer.Option{
		producer.WithCredentials(credentials),
		producer.WithGroupName(config.GroupName),
		producer.WithVIPChannel(config.VipChannelEnabled),
		producer.WithQueueSelector(NewSelector(int(config.Selector))),
	}

	if err := primitive.NamesrvAddr(config.Namesrvs).Check(); err != nil {
		ops = append(ops, producer.WithNameServerDomain(config.Namesrvs[0]))
	} else {
		ops = append(ops, producer.WithNameServer(config.Namesrvs))
	}

	if config.InstanceName != "" {
		ops = append(ops, producer.WithInstanceName(config.InstanceName))
	} else {
		ops = append(ops, producer.WithInstanceName(strings.Join(config.Namesrvs, ";")))
	}

	if config.Namespace != "" {
		ops = append(ops, producer.WithNamespace(config.Namespace))
	}

	if config.SendMsgTimeout > 0 {
		ops = append(ops, producer.WithSendMsgTimeout(config.SendMsgTimeout*time.Millisecond))
	}

	if config.RetryTimes > 0 {
		ops = append(ops, producer.WithRetry(config.RetryTimes))
	}

	if config.TopicQueueNums > 0 {
		ops = append(ops, producer.WithDefaultTopicQueueNums(config.TopicQueueNums))
	}

	if config.TopicKey != "" {
		ops = append(ops, producer.WithCreateTopicKey(config.TopicKey))
	}

	return
}

// ProducerOptionsWithACM 根据 ACM 配置信息构造生产者参数
func ProducerOptionsWithACM(section string) []producer.Option {
	buildKey := func(key string) string {
		return strings.Join([]string{section, key}, ".")
	}

	config := ProducerConfig{
		Namesrvs: acm.GetStringSlice(buildKey("namesrvs")),
		Auth: Auth{
			AK:    acm.GetString(buildKey("auth.ak")),
			SK:    acm.GetString(buildKey("auth.sk")),
			Token: acm.GetString(buildKey("auth.token")),
		},
		TopicName:         acm.GetString(buildKey("topic_name")),
		GroupName:         acm.GetString(buildKey("group_name")),
		Namespace:         acm.GetString(buildKey("namespace")),
		InstanceName:      acm.GetString(buildKey("instance_name")),
		RetryTimes:        acm.GetInt(buildKey("retry_times")),
		Selector:          acm.GetInt(buildKey("selector")),
		SendMsgTimeout:    acm.GetDuration(buildKey("send_msg_timeout")),
		TopicKey:          acm.GetString(buildKey("topic_key")),
		TopicQueueNums:    acm.GetInt(buildKey("topic_queue_nums")),
		VipChannelEnabled: acm.GetBool(buildKey("vip_channel_enabled")),
	}
	return ProducerOptions(config)
}
