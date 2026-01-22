package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.Logger
}

func NewLogger(level string) (*Logger, error) {
	var cfg zap.Config
	if level == "debug" {
		cfg = zap.NewDevelopmentConfig()
	} else {
		cfg = zap.NewProductionConfig()
	}

	// Parse log level
	l, err := zapcore.ParseLevel(level)
	if err != nil {
		l = zap.InfoLevel
	}
	cfg.Level = zap.NewAtomicLevelAt(l)

	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	return &Logger{logger}, nil
}
