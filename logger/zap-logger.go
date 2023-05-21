package logger

import (
	"fmt"

	config "github.com/scripvoice/core/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	logger *zap.Logger
}

func NewZapLogger(zapConfig config.ZapConfig) (ILogger, error) {
	//config := zap.NewProductionConfig()
	//config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	// Set the logging level
	level := zap.NewAtomicLevel()
	err := level.UnmarshalText([]byte(zapConfig.Level))
	if err != nil {
		fmt.Println("Failed to parse log level:", err)
		return nil, err
	}

	// Create the Zap logger config
	zapConfigStruct := zap.Config{
		Level:            level,
		Encoding:         zapConfig.Encoding,
		OutputPaths:      zapConfig.OutputPaths,
		ErrorOutputPaths: zapConfig.ErrorOutputPaths,
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",
			LevelKey:   "level",
		},
	}

	// Build the logger
	logger, err := zapConfigStruct.Build()
	if err != nil {
		fmt.Println("Failed to build logger:", err)
		return nil, err
	}

	fmt.Println("logger initialized successfully.")
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
