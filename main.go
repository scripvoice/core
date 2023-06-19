package main

import (
	"context"

	"github.com/scripvoice/core/logger"
	"golang.org/x/exp/slog"
)

func main() {
	testLogger()
}

func testLogger() {
	config := logger.LoggerConfig{
		Level:   slog.LevelInfo,
		Service: "CoreService",
		Handler: logger.HandlerOptions{

			Sink:             "Postgres",
			ConnectionString: "postgres://postgres:naval@127.0.0.1:5433/test?sslmode=disable",

			Levels: nil,
			HandlerOptions: slog.HandlerOptions{
				AddSource: true,
			},
		},
	}
	logger.InitLogger(config, context.Background())

	// handler := slog.NewJSONHandler(os.Stdout, nil)
	// log := slog.New(handler)
	// slog.SetDefault(log)

}
