package rocket

import (
	"sync"

	log "MyIM/pkg/mlog"
	"MyIM/pkg/rocketmq-client-go/rlog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger Logger
type Logger struct {
	*log.Mlog
	level log.Level
}

var defaultLogger *Logger
var defaultLoggerOnce sync.Once

// InitLogger 初始化日志
func InitLogger(Mlog *log.Mlog) {
	defaultLoggerOnce.Do(func() {
		defaultLogger = newLogger(Mlog, log.WarnLevel)
		rlog.SetLogger(defaultLogger)
	})
}

func newLogger(Mlog *log.Mlog, level log.Level) *Logger {
	if level == "" {
		level = log.WarnLevel
	}

	return &Logger{
		Mlog:  Mlog,
		level: level,
	}
}

// OutputPath log file path
// Notice: 不实现,使用应用日志文件
func (l *Logger) OutputPath(path string) (err error) {
	return nil
}

// Level Level
// Notice: 不实现,不支持动态修改
func (l *Logger) Level(level string) {}

// Debug debug level
func (l *Logger) Debug(msg string, fields map[string]interface{}) {
	if !log.Enabled(l.level, zapcore.DebugLevel) {
		return
	}
	l.Logger.Debug(msg, l.getZapFields(fields)...)
}

// Info info level
func (l *Logger) Info(msg string, fields map[string]interface{}) {
	if !log.Enabled(l.level, zapcore.InfoLevel) {
		return
	}
	l.Logger.Info(msg, l.getZapFields(fields)...)
}

// Warning warn level
func (l *Logger) Warning(msg string, fields map[string]interface{}) {
	if !log.Enabled(l.level, zapcore.WarnLevel) {
		return
	}
	l.Logger.Warn(msg, l.getZapFields(fields)...)
}

// Error error level
func (l *Logger) Error(msg string, fields map[string]interface{}) {
	if !log.Enabled(l.level, zapcore.ErrorLevel) {
		return
	}
	l.Logger.Error(msg, l.getZapFields(fields)...)
}

// Fatal fatal level
func (l *Logger) Fatal(msg string, fields map[string]interface{}) {
	if !log.Enabled(l.level, zapcore.ErrorLevel) {
		return
	}
	l.Logger.Error(msg, l.getZapFields(fields)...)
}

func (l *Logger) getZapFields(fields map[string]interface{}) (zapFields []zap.Field) {
	zapFields = []zap.Field{
		zap.String(log.NameKey, "rocket"),
	}
	for k, v := range fields {
		zapFields = append(zapFields, zap.Any(k, v))
	}
	return
}
