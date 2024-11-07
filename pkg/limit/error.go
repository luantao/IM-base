package limit

import (
	sentinelbase "github.com/alibaba/sentinel-golang/core/base"
	"github.com/luantao/IM-base/pkg/rediss"
	"time"
)

// IsLimit 是否触发限流
func IsLimit(err error) bool {
	if _, ok := err.(*sentinelbase.BlockError); ok {
		return true
	}

	if _, ok := err.(*ClusterLimitError); ok {
		return true
	}
	return false
}

// ClusterLimitError redis 集群限流错误类型
type ClusterLimitError struct {
	ErrMsg string
	Result *rediss.Result
}

func (e ClusterLimitError) Error() string {
	return e.ErrMsg
}

type Limit struct {
	Rate   int
	Burst  int
	Period time.Duration
}
