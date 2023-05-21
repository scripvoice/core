package logger

import (
	"fmt"
	"sync"

	config "github.com/scripvoice/core/config"
)

var (
	loggerInstance ILogger
	oncelogger     sync.Once
)

// GetInstance returns the singleton instance
func InitLogger(config config.ZapConfig) (ILogger, error) {
	var err error
	oncelogger.Do(func() {
		loggerInstance, err = NewZapLogger(config) // Create the singleton instance
	})
	if err != nil {
		fmt.Println(err)
		panic("error creating logger")
	}
	return loggerInstance, err
}

func GetLogger() (ILogger, error) {
	if loggerInstance == nil {

		return nil, nil
	}
	return loggerInstance, nil

}
