package log

import (
	"go.uber.org/zap"
)

// Global variables against which all our logging occurs.
var logger = zap.NewExample()
var sugar = logger.Sugar()

// https://pkg.go.dev/go.uber.org/zap#SugaredLogger.Error
func Error(args ...any) error {
	sugar.Error(args...)
	return nil
}

// https://pkg.go.dev/go.uber.org/zap#SugaredLogger.Error
func Fatal(args ...any) error {
	sugar.Fatal(args...)
	return nil
}
