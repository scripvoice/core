package logger

import (
	"fmt"
	"sync"

	"go.uber.org/zap"
)

var (
	loggerInstance ILogger
	oncelogger     sync.Once
)

// GetInstance returns the singleton instance
func GetLoggerInstance(config zap.Config) (ILogger, error) {
	var err error
	oncelogger.Do(func() {
		loggerInstance, err = NewZapLogger(config) // Create the singleton instance
	})
	if err != nil {
		fmt.Println("error creating logger")
		panic("error creating logger")
	}
	return loggerInstance, err
}
