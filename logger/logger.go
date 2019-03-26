package logger

import (
	"go.uber.org/zap"
)

type Logger struct {
}

func Initialize() {
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
}

func Infow(msg string, args ...interface{}) {
	zap.S().Infow(msg, args...)
}

func Infof(template string, args ...interface{}) {
	zap.S().Infof(template, args...)
}
