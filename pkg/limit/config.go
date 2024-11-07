package limit

// 项目类型
const (
	AppTypeDefault = 0 // 默认,不区分项目类型
	AppTypeWeb     = 1 // web service
	AppTypeRPC     = 2 // RPC service
	AppTypeDaemon  = 3 // daemon service
)

// 默认配置
const (
	DefaultMetricMaxFileCount    = 1
	DefaultMetricMaxFileSize     = 1024 * 1024 * 1024 // 1G
	DefaultMetricFlushInterval   = 1                  // 1 秒;单位:秒
	DefaultSystemCollectInterval = 1000               // 1 秒;单位:毫秒
)

// Config 配置信息
//
// yaml example:
// sentinel:
//
//	app_name: xxx-order-center
//	log_path: ./logs/
//	log_level: debug # debug/info/warn/error
//	metric_max_file_count: 1 # 最多文件数;默认 1
//	metric_max_file_size: 1024 # 单个文件大小;单位 M;默认 1G
//	metric_flush_interval: 1 # 聚合刷盘间隔时间;单位:秒;默认 1 秒
//	system_collect_interval: 1000 # 指标收集间隔时间;单位:毫秒;默认 1 秒
type Config struct {
	AppName               string  `json:"app_name" yaml:"app_name"`   // 项目名字
	LogPath               string  `json:"log_path" yaml:"log_path"`   // 日志文件目录
	LogLevel              string  `json:"log_level" yaml:"log_level"` // debug/info/warn/error
	MetricMaxFileCount    uint32  `json:"metric_max_file_count" yaml:"metric_max_file_count"`
	MetricMaxFileSize     uint64  `json:"metric_max_file_size" yaml:"metric_max_file_size"`
	MetricFlushInterval   uint32  `json:"metric_flush_interval" yaml:"metric_flush_interval"`
	SystemCollectInterval uint32  `json:"system_collect_interval" yaml:"system_collect_interval"`
	Logger                *Logger `json:"-" yaml:"-"`
}
