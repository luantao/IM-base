package rocket

import (
	"MyIM/pkg/rocketmq-client-go/producer"
)

// NewSelector 选择分区算法
//
// @param selectorType 分区算法类型
// - 0: roundRobin
// - 1: hash
// - 2: random
// - 3: manual
// - other: roundRobin
func NewSelector(selectorType int) producer.QueueSelector {
	switch selectorType {
	case 0:
		return producer.NewRoundRobinQueueSelector()
	case 1:
		return producer.NewHashQueueSelector()
	case 2:
		return producer.NewRandomQueueSelector()
	case 3:
		return producer.NewManualQueueSelector()
	default:
		return producer.NewRoundRobinQueueSelector()
	}
}
