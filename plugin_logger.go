package logger

import (
	"github.com/edwingeng/slog"
)

var PluginLogger PluginLoggerSwapper

type PluginLoggerSwapper struct{}

func (plog PluginLoggerSwapper) NewLoggerWith(keyVals ...interface{}) slog.Logger {
	return plog
}

func (plog PluginLoggerSwapper) Debug(args ...interface{}) {
	DebugDepth(1, args...)
}

func (plog PluginLoggerSwapper) Info(args ...interface{}) {
	InfoDepth(1, args...)
}

func (plog PluginLoggerSwapper) Warn(args ...interface{}) {
	WarningDepth(1, args...)
}

func (plog PluginLoggerSwapper) Error(args ...interface{}) {
	ErrorDepth(1, args...)
}

func (plog PluginLoggerSwapper) Debugf(template string, args ...interface{}) {
	DebugfDepth(1, template, args...)
}

func (plog PluginLoggerSwapper) Infof(template string, args ...interface{}) {
	InfofDepth(1, template, args...)
}

func (plog PluginLoggerSwapper) Warnf(template string, args ...interface{}) {
	WarningfDepth(1, template, args...)
}

func (plog PluginLoggerSwapper) Errorf(template string, args ...interface{}) {
	ErrorfDepth(1, template, args...)
}

func (plog PluginLoggerSwapper) Debugw(msg string, keyVals ...interface{}) {
	DebugfDepth(1, msg, keyVals...)
}

func (plog PluginLoggerSwapper) Infow(msg string, keyVals ...interface{}) {
	InfofDepth(1, msg, keyVals...)
}

func (plog PluginLoggerSwapper) Warnw(msg string, keyVals ...interface{}) {
	WarningfDepth(1, msg, keyVals...)
}

func (plog PluginLoggerSwapper) Errorw(msg string, keyVals ...interface{}) {
	ErrorfDepth(1, msg, keyVals...)
}

func (plog PluginLoggerSwapper) FlushLogger() error {
	logging.lockAndFlushAll()
	return nil
}
