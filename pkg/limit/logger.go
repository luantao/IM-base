package limit

import (
	"MyIM/pkg/mlog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger Logger
type Logger struct {
	*mlog.Mlog
	level mlog.Level
}

// NewLogger new logger.
func NewLogger(Mlog *mlog.Mlog, level mlog.Level) *Logger {
	if level == "" {
		level = mlog.WarnLevel
	}

	return &Logger{
		Mlog:  Mlog,
		level: level,
	}
}

func (l Logger) DebugEnabled() bool {
	return mlog.Enabled(l.level, zapcore.DebugLevel)
}

func (l Logger) Debug(msg string, args ...interface{}) {
	l.Mlog.Debug(msg, sweetenFields(args)...)
}

func (l Logger) InfoEnabled() bool {
	return mlog.Enabled(l.level, zapcore.InfoLevel)
}

func (l Logger) Info(msg string, args ...interface{}) {
	l.Mlog.Info(msg, sweetenFields(args)...)
}

func (l Logger) WarnEnabled() bool {
	return mlog.Enabled(l.level, zapcore.WarnLevel)
}

func (l Logger) Warn(msg string, args ...interface{}) {
	l.Mlog.Warn(msg, sweetenFields(args)...)
}

func (l Logger) ErrorEnabled() bool {
	return mlog.Enabled(l.level, zapcore.ErrorLevel)
}

func (l Logger) Error(err error, msg string, args ...interface{}) {
	fields := append(sweetenFields(args), zap.Error(err))
	l.Mlog.Error(msg, fields...)
}

func sweetenFields(args []interface{}) []zap.Field {
	fields := []zap.Field{
		zap.String(mlog.NameKey, "sentinel"),
	}

	var key interface{}
	for i, v := range args {
		if i%2 != 0 {
			key = v
			continue
		}

		keyStr, ok := key.(string)
		if !ok {
			continue
		}
		fields = append(fields, zap.Any(keyStr, v))
	}

	return fields
}
