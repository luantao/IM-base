package http

import (
	"MyIM/pkg/mlog"
	"context"
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger Logger
type Logger struct {
	*mlog.Mlog
	level mlog.Level
	ctx   context.Context // 构建 resty log 使用，获取 traceID
}

// NewLogger new logger.
func NewLogger(zlog *mlog.Mlog, level mlog.Level) *Logger {
	if level == "" {
		level = mlog.WarnLevel
	}

	return &Logger{
		Mlog:  zlog,
		level: level,
	}
}

// RestyLogger 构建 go-resty 日志
// @param ctx 获取 traceID
func (l Logger) RestyLogger(ctx context.Context) *Logger {
	l.ctx = ctx
	return &l
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	if !mlog.Enabled(l.level, zapcore.ErrorLevel) {
		return
	}
	l.WithCTX(l.ctx).Error(fmt.Sprintf(format, v...), zap.String(mlog.NameKey, "resty"))
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	if !mlog.Enabled(l.level, zapcore.WarnLevel) {
		return
	}
	l.WithCTX(l.ctx).Warn(fmt.Sprintf(format, v...), zap.String(mlog.NameKey, "resty"))
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	if !mlog.Enabled(l.level, zapcore.DebugLevel) {
		return
	}
	l.WithCTX(l.ctx).Debug(fmt.Sprintf(format, v...), zap.String(mlog.NameKey, "resty"))
}

// no-op resty logger
type nopRestyLogger struct{}

func NewNopRestyLogger() nopRestyLogger                       { return nopRestyLogger{} }
func (nopRestyLogger) Errorf(format string, v ...interface{}) {}
func (nopRestyLogger) Warnf(format string, v ...interface{})  {}
func (nopRestyLogger) Debugf(format string, v ...interface{}) {}
