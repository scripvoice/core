package logger

import (
	"go.uber.org/zap"
)

type ZapLogger struct {
	logger *zap.Logger
}

func NewZapLogger(config zap.Config) (ILogger, error) {
	//config := zap.NewProductionConfig()
	//config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, err := config.Build()
	if err != nil {
		return nil, err
	}
	return ZapLogger{
		logger: logger,
	}, nil
}

func (l ZapLogger) Info(message string, fields ...interface{}) {
	l.logger.Info(message, convertFields(fields)...)
}

func (l ZapLogger) Error(message string, fields ...interface{}) {
	l.logger.Error(message, convertFields(fields)...)
}

func (l ZapLogger) Debug(message string, fields ...interface{}) {
	l.logger.Debug(message, convertFields(fields)...)
}

func (l ZapLogger) Sync() error {
	return l.logger.Sync()
}

func convertFields(fields []interface{}) []zap.Field {
	zapFields := make([]zap.Field, len(fields))
	for i, f := range fields {
		zapFields[i] = zap.Any("field", f)
	}
	return zapFields
}
