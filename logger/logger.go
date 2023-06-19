package logger

import (
	"context"
	"sync"

	"golang.org/x/exp/slices"
	"golang.org/x/exp/slog"
)

var Config LoggerConfig

type LoggerConfig struct {
	Level   slog.Level
	Service string
	Handler HandlerOptions
}

type HandlerOptions struct {
	Sink             string
	ConnectionString string
	Levels           []slog.Level
	slog.HandlerOptions
}

type LogMessage struct {
	Message string
	Mu      sync.Mutex
}

func (m *LogMessage) Write(p []byte) (n int, err error) {
	m.Message = string(p)
	return 1, nil
}

func GetRquestAttributeFromContext(c context.Context) (userId string, TraceId string) {
	uId := c.Value("UserId")
	if uId != nil {
		if v, ok := uId.(string); ok {
			userId = v
		}
	}

	tId := c.Value("TraceId")
	if tId != nil {
		if v, ok := tId.(string); ok {
			TraceId = v
		}
	}
	return
}

func InitLogger(config LoggerConfig, ctx context.Context) {

	Config = config
	if config.Handler.Sink == "Postgres" {
		logger := slog.New(NewPostgresqlHandler(context.Background(), config.Handler)).With(slog.String("Service", config.Service))
		slog.SetDefault(logger)
	}
}

type CustomHandler struct {
	HandlerOptions
}

func (h *CustomHandler) Enabled(c context.Context, level slog.Level) bool {
	switch {
	case len(h.Levels) > 0 && slices.Contains(h.Levels, level):
		return true
	case (h.Level != nil && h.Level.Level() >= level) || Config.Level >= level:
		return true
	default:
		return false
	}
}

type Logger interface {
	Log(ctx context.Context, level slog.Level, msg string, args ...any)
	LogAttrs(ctx context.Context, level slog.Level, msg string, attrs ...slog.Attr)
	DebugCtx(ctx context.Context, msg string, args ...any)
	InfoCtx(ctx context.Context, msg string, args ...any)
	WarnCtx(ctx context.Context, msg string, args ...any)
	ErrorCtx(ctx context.Context, msg string, args ...any)
}

/* Call this method to get the logger*/
func GetLogger() Logger {
	return slog.Default()
}
