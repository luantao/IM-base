package mlog

const (
	// 日志类型
	LogTypeMQ      = "mq"
	LogTypeAPI     = "api"
	LogTypeBiz     = "biz"
	LogTypeEvent   = "track"
	LogTypeMonitor = "monitor"

	// 日志字段名
	LogEventKey   = "log_event"    // 其它日志标识;非数仓埋点
	LogTypeKey    = "log_type"     // 日志类型
	ContentKey    = "properties"   // 日志索引字段
	APIRespKey    = "response"     // API 返回值索引
	APIReqBodyKey = "request_body" // 请求BODY
)
