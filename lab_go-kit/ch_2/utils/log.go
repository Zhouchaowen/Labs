package utils

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func NewLogger(service string) (*zap.SugaredLogger, error) {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout"}
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.DisableStacktrace = true
	config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	config.InitialFields = map[string]interface{}{
		"service": service,
	}

	var err error
	logger, err = config.Build()
	if err != nil {
		return nil, err
	}

	return logger.Sugar(), nil
}

func GetLogger() *zap.Logger {
	return logger
}
