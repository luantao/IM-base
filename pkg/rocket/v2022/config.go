package rocket

import "time"

type Auth struct {
	AK    string `json:"ak" yaml:"ak"`       // AccessKey
	SK    string `json:"sk" yaml:"sk"`       // SecretKey
	Token string `json:"token" yaml:"token"` // SecurityToken
}

// ProducerConfig 生产者配置信息
//
// 被 # 注释的选项为选填
// example_producer:
//
//	namesrvs:
//	  - 172.17.12.48:9876
//	# auth:
//	#   ak: xxx
//	#   sk: xxx
//	#   token: xxx
//	topic_name: topic_oc_example
//	group_name: producer_group_oc_example
//	# namespace: test_namespace
//	# instance_name: test_instance # 默认 PID
//	# retry_times: 3 # 重试次数
//	# selector: 0 # 分区算法;1:hash;2:random;3:manual;other:roundRobin
//	# send_msg_timeout: 3000 # 单位:毫秒
//	# topic_key: example_topic_key
//	# topic_queue_nums: 4 #  队列数量
//	# vip_channel_enabled: false
type ProducerConfig struct {
	Namesrvs          []string      `json:"namesrvs" yaml:"namesrvs"`
	Auth              Auth          `json:"auth" yaml:"auth"`
	TopicName         string        `json:"topic_name" yaml:"topic_name"`
	GroupName         string        `json:"group_name" yaml:"group_name"`
	Namespace         string        `json:"namespace" yaml:"namespace"`
	InstanceName      string        `json:"instance_name" yaml:"instance_name"`
	RetryTimes        int           `json:"retry_times" yaml:"retry_times"`
	Selector          int           `json:"selector" yaml:"selector"`
	SendMsgTimeout    time.Duration `json:"send_msg_timeout" yaml:"send_msg_timeout"`
	TopicKey          string        `json:"topic_key" yaml:"topic_key"`
	TopicQueueNums    int           `json:"topic_queue_nums" yaml:"topic_queue_nums"`
	VipChannelEnabled bool          `json:"vip_channel_enabled" yaml:"vip_channel_enabled"`
}

// ConsumerConfig 消费者配置信息
//
// 被 # 注释的选项为选填
// example_consumer:
//
//	namesrvs:
//	  - 127.0.0.1:9876
//	# auth:
//	#   ak: xxx
//	#   sk: xxx
//	#   token: xxx
//	topic_name: topic_oc_example
//	group_name: consumer_group_oc_example
//	# namespace: test_namespace
//	# instance_name: test_instance # 默认 PID
//	# tags: "" # consumer tags "tag1 || tag2"
//	consumer_model: 1 # 0:广播模式;1:集群模式
//	# consumer_from_where: 0 # 启动时的消费位点;0:最后一条;1:第一条;2:时间
//	# consumer_orderly: false # 是否顺序消费
//	# max_reconsume_times: -1 # 最大重试消费次数
//	# pull_batch_size: 32
//	# message_batch_max_size: 512
//	auto_commit: true
//	# rebalance_lock_interval: 20000 # 单位:毫秒
//	# suspend_current_queue_time: 1000 # 单位:毫秒
//	# pull_interval: 100 # 单位:毫秒
//	# vip_channel_enabled: false
//	# max_reconsume_times: -1 # 最大重试消费次数
type ConsumerConfig struct {
	Namesrvs                []string      `json:"namesrvs" yaml:"namesrvs"`
	Auth                    Auth          `json:"auth" yaml:"auth"`
	TopicName               string        `json:"topic_name" yaml:"topic_name"`
	GroupName               string        `json:"group_name" yaml:"group_name"`
	Namespace               string        `json:"namespace" yaml:"namespace"`
	InstanceName            string        `json:"instance_name" yaml:"instance_name"`
	Tags                    string        `json:"tags" yaml:"tags"`
	ConsumerModel           int           `json:"consumer_model" yaml:"consumer_model"`
	ConsumerFromWhere       int           `json:"consumer_from_where" yaml:"consumer_from_where"`
	ConsumerOrderly         bool          `json:"consumer_orderly" yaml:"consumer_orderly"`
	MaxReconsumeTimes       int32         `json:"max_reconsume_times" yaml:"max_reconsume_times"`
	PullBatchSize           int32         `json:"pull_batch_size" yaml:"pull_batch_size"`
	MessageBatchMaxSize     int           `json:"message_batch_max_size" yaml:"message_batch_max_size"`
	AutoCommit              bool          `json:"auto_commit" yaml:"auto_commit"`
	RebalanceLockInterval   time.Duration `json:"rebalance_lock_interval" yaml:"rebalance_lock_interval"`
	SuspendCurrentQueueTime time.Duration `json:"suspend_current_queue_time" yaml:"suspend_current_queue_time"`
	PullInterval            time.Duration `json:"pull_interval" yaml:"pull_interval"`
	VipChannelEnabled       bool          `json:"vip_channel_enabled" yaml:"vip_channel_enabled"`
}
