package mlog

import (
	"context"
	"fmt"
	"github.com/luantao/IM-base/pkg/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

type Level = string

// gobase 各 package 输出的日志级别
const (
	DisableLevel Level = "disable" // 禁用（不打日志）
	DebugLevel   Level = "debug"
	InfoLevel    Level = "info"
	WarnLevel    Level = "warn"
	ErrorLevel   Level = "error"
)
const (
	TraceID   = "trace_id"   // 链路追踪 ID
	RequestID = "request_id" // 请求唯一 ID
	StartTime = "start_time" // 请求开始时间
	EventKey  = "event"      // 数仓埋点日志标识
)

const NameKey = "category"

// Mlog 标准日志
type Mlog struct {
	*zap.Logger
	level zap.AtomicLevel
}

func (l Mlog) Print(v ...interface{}) {
	l.Info("DB", zap.Any("Mysql", v))
}
func (l Mlog) Printf(str string, fields ...interface{}) {
	if _, ok := fields[1].(error); ok {
		// err 是一个error类型
		l.SetError(fmt.Sprintf(str, fields...), fields[1].(error), zap.Any("DB", "MySQL"))
	} else {
		l.Info(fmt.Sprintf(str, fields...), zap.Any("DB", "MySQL"))
	}
}

// SetLevel 设置日志级别
func (l Mlog) SetLevel(levelName Level) {
	level := GetLevelType(levelName)
	if level == l.GetLevel() {
		return
	}
	l.level.SetLevel(level)
}

// GetLevel 获取当前日志级别
func (l Mlog) GetLevel() zapcore.Level {
	return l.level.Level()
}

// Named 日志分类 (封装zap.Logger.Named)
func (l Mlog) Named(name string) Mlog {
	l.Logger = l.Logger.Named(name)
	return l
}

// WithCTX 从上下文中获取 traceid 并在日志中加入 traceid 字段
func (l Mlog) WithCTX(c context.Context) Mlog {
	if c == nil {
		return l
	}
	var fields []zap.Field
	traceID, ok := c.Value(TraceID).(string)
	if ok {
		fields = []zap.Field{
			zap.String(TraceID, traceID),
		}
	}

	requestID, ok := c.Value(RequestID).(string)
	if ok {
		fields = append(fields, zap.String(RequestID, requestID))
	}

	if len(fields) == 0 {
		return l
	}

	l.Logger = l.With(fields...)
	return l
}
func (l Mlog) SetError(msg string, err error, fields ...zap.Field) Mlog {
	if err == nil {
		return l
	}
	errStr := "\u001B[1;30;41m[" + err.Error() + "]\u001B[0m "
	if config.GetString("app.env") != "release" {
		fmt.Println(errStr)
	}
	fields = append(fields, zap.Error(err))
	l.Error(msg, fields...)
	return l
}

// WithEvent 数据埋点，只有 event 值存在时，日志方才投递到大数据中心
func (l Mlog) WithEvent(eventName string) Mlog {
	l.Logger = l.With(
		zap.String(EventKey, eventName),
	)
	return l
}

// GetLevelType 获取日志级别类型
func GetLevelType(levelName Level) zapcore.Level {
	var l zapcore.Level
	switch levelName {
	case DebugLevel:
		l = zap.DebugLevel
	case InfoLevel:
		l = zap.InfoLevel
	case WarnLevel:
		l = zap.WarnLevel
	case ErrorLevel:
		l = zap.ErrorLevel
	default:
		l = zap.InfoLevel
	}
	return l
}

func newZapCore(cfg Config, level zap.AtomicLevel) zapcore.Core {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        NameKey,
		CallerKey:      "line",
		MessageKey:     "_msg",
		StacktraceKey:  "stack",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 所有日志级别使用同一个日志文件
	if !cfg.LevelFile {
		return zapcore.NewTee(
			zapcore.NewCore(
				zapcore.NewJSONEncoder(encoderConfig),
				zapcore.NewMultiWriteSyncer(writers(cfg, "")...),
				level),
		)
	}

	// 每个日志级别使用单独的文件
	var cores []zapcore.Core
	var l zapcore.Level
	for l = zapcore.DebugLevel; l <= zapcore.FatalLevel; l++ {
		coreLevel := l
		cores = append(cores,
			zapcore.NewCore(
				zapcore.NewJSONEncoder(encoderConfig),
				zapcore.NewMultiWriteSyncer(writers(cfg, l.String())...),
				zap.LevelEnablerFunc(func(itemLevel zapcore.Level) bool {
					return coreLevel >= level.Level() && itemLevel == coreLevel
				})),
		)
	}
	return zapcore.NewTee(cores...)
}

// writers 日志输出
func writers(cfg Config, level string) (ws []zapcore.WriteSyncer) {
	var (
		levelPath string // e.g. -info
		filePath  string // e.g. /tmp/app-info.log
		handle    lumberjack.Logger
	)

	if level != "" {
		levelPath = "-" + level
	}
	filePath = cfg.LogPath + cfg.LogFileName + levelPath + ".log"

	handle = lumberjack.Logger{
		Filename:   filePath,
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
		Compress:   cfg.Compress,
		LocalTime:  true,
	}
	ws = []zapcore.WriteSyncer{
		zapcore.AddSync(&handle),
	}
	if cfg.Stdout {
		ws = append(ws, zapcore.AddSync(os.Stdout))
	}
	return
}

// Enabled 是否输出日志
// @param configLevel 预期设置的日志输出级别
// @param outLevel 正在打印的日志输出级别
func Enabled(configLevel Level, outLevel zapcore.Level) bool {
	if configLevel == DisableLevel {
		return false
	}
	return GetLevelType(configLevel) <= outLevel
}
