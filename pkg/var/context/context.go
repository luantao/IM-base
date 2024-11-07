package context

// 上下文中自定义变量
const (
	TraceID     = "trace_id"   // 链路追踪 ID
	RequestID   = "request_id" // 请求唯一 ID
	StartTime   = "start_time" // 请求开始时间
	AppID       = "app_id"     // api 认证信息
	Environment = "env"        // 环境标识
	Version     = "ver"        //版本标识
)
