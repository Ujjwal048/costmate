package logger

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	// Logger is the global logger instance
	Logger *zap.Logger
)

// Initialize sets up the logger
func Initialize() error {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeCaller = nil 
	config.OutputPaths = []string{"/tmp/costmate.log"}
	config.ErrorOutputPaths = []string{"/tmp/costmate.log"}
	config.Encoding = "console"
	config.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Local().Format("02-01-2006T15:04:05"))
	}

	var err error
	Logger, err = config.Build()
	if err != nil {
		return err
	}

	return nil
}

// Close closes the logger
func Close() {
	if Logger != nil {
		Logger.Sync()
	}
}

// Fatal logs a fatal message and exits
func Fatal(msg string, err error) {
	if Logger != nil {
		Logger.Fatal(msg, zap.Error(err))
	} else {
		os.Exit(1)
	}
}

// Error logs an error message
func Error(msg string, err error) {
	if Logger != nil {
		Logger.Error(msg, zap.Error(err))
	}
}

// Info logs an info message
func Info(msg string, args ...interface{}) {
	if Logger != nil {
		Logger.Sugar().Infof(msg, args...)
	}
}

// Debug logs a debug message
func Debug(msg string, args ...interface{}) {
	if Logger != nil {
		Logger.Sugar().Debugf(msg, args...)
	}
}
