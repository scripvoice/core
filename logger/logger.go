package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var DefaultConfig = zap.Config{
	Encoding:         "json",
	OutputPaths:      []string{"stdout", "log.log"},
	ErrorOutputPaths: []string{"stderr", "log.log"},
	EncoderConfig: zapcore.EncoderConfig{
		MessageKey: "message",
		LevelKey:   "level",
	},
}
