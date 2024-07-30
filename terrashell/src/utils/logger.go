package utils

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/logger"
)

type Logger struct {
	Context *testing.T
}

var instanceLogger *Logger

func LoggerGetInstance() *Logger {
	if instanceLogger == nil {
		instanceLogger = &Logger{
			Context: &testing.T{},
		}
	}
	return instanceLogger
}

func (l *Logger) Log(args ...interface{}) {
	logger.Log(l.Context, args)
}
