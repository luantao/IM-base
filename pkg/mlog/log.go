package mlog

import (
	"MyIM/pkg/config"
	"context"
	"strings"
	"sync"

	"go.uber.org/zap"
)

// Config 日志配置参数
//
// yaml example:
// log:
//
//	app_name: xx-order-center # 项目名字
//	log_path: ./logs/ # 日志文件目录
//	log_file_name: app # 日志文件名字前缀
//	log_level: error # debug/info/warn/error
//	level_file: false # 是否启用单独级别对应单独文件
//	stdout: false # 是否打印到标准输出
//	max_age: 2 # 备份文件保留时间;单位:天
//	max_size: 1024 # 单个日志文件大小;单位:M
//	max_backups: 2 # 最大备份日志文件数量
//	compress: false # 是否启用压缩
type Config struct {
	AppName     string `json:"app_name" yaml:"app_name"`           // 项目名字
	LogPath     string `json:"log_path" yaml:"log_path"`           // 日志文件目录
	LogFileName string `json:"log_file_name" yaml:"log_file_name"` // 日志文件名字前缀
	LogLevel    string `json:"log_level" yaml:"log_level"`         // debug/info/warn/error
	LevelFile   bool   `json:"level_file" yaml:"level_file"`       // 是否启用单独级别对应单独文件
	Stdout      bool   `json:"stdout" yaml:"stdout"`               // 是否打印到标准输出
	MaxAge      int    `json:"max_age" yaml:"max_age"`             // 备份文件保留时间;单位:天
	MaxSize     int    `json:"max_size" yaml:"max_size"`           // 单个日志文件大小;单位:M
	MaxBackups  int    `json:"max_backups" yaml:"max_backups"`     // 最大备份日志文件数量
	Compress    bool   `json:"compress" yaml:"compress"`           // 是否启用压缩
}

// 默认日志对象
var defaultLogger *Mlog
var defaultOnce sync.Once

// Init init default Logger.
// 从 Config 中获取配置初始化默认日志对象
// @param section 日志配置信息所属节点
func Init(section string) {
	defaultOnce.Do(func() {
		defaultLogger = NewWithACM(section)
	})
}

// Logger return default Logger.
func Logger() *Mlog {
	return defaultLogger
}

// LoggerWithCTX return Logger with Context.
func LoggerWithCTX(ctx context.Context) Mlog {
	return defaultLogger.WithCTX(ctx)
}

// New return a new logger with config.
func New(cfg Config) *Mlog {
	level := zap.NewAtomicLevel()

	zapLogger := zap.New(
		newZapCore(cfg, level),
		zap.AddCaller(),
		// zap.Development(),
	).
		With(
			zap.String("app_name", cfg.AppName),
		)

	zlog := &Mlog{
		Logger: zapLogger,
		level:  level,
	}
	zlog.SetLevel(cfg.LogLevel)
	return zlog
}

// NewWithACM return a new logger from Config.
func NewWithACM(section string) *Mlog {
	buildKey := func(key string) string {
		return strings.Join([]string{section, key}, ".")
	}

	cfg := Config{
		AppName:     config.GetString(buildKey("app_name")),
		LogPath:     config.GetString(buildKey("log_path")),
		LogFileName: config.GetString(buildKey("log_file_name")),
		LogLevel:    config.GetString(buildKey("log_level")),
		LevelFile:   config.GetBool(buildKey("level_file")),
		Stdout:      config.GetBool(buildKey("stdout")),
		MaxAge:      config.GetInt(buildKey("max_age")),
		MaxSize:     config.GetInt(buildKey("max_size")),
		MaxBackups:  config.GetInt(buildKey("max_backups")),
		Compress:    config.GetBool(buildKey("compress")),
	}
	return New(cfg)
}

// NewExample builds a Logger that's designed for use in zap's testable
// examples. It writes DebugLevel and above logs to standard out as JSON, but
// omits the timestamp and calling function to keep example output
// short and deterministic.
func NewExample() *Mlog {
	level := zap.NewAtomicLevelAt(zap.DebugLevel)
	zapLogger := zap.NewExample()
	zlog := &Mlog{
		Logger: zapLogger,
		level:  level,
	}
	return zlog
}
