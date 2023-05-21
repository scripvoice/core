package logger

// ILogger is the interface for a logger
type ILogger interface {
	Info(message string, fields ...interface{})
	Error(message string, fields ...interface{})
	Debug(message string, fields ...interface{})
	Sync() error
}
