package logger

import (
	"context"
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
	"golang.org/x/exp/slog"
)

type PostgresqlHandler struct {
	CustomHandler
	JSONHandler *slog.JSONHandler
	Ctx         context.Context
}

var db *sql.DB
var dbInitialize sync.Once
var message LogMessage = LogMessage{}

func (h *PostgresqlHandler) WriteToDb(message string, ctx context.Context) error {

	con, err := db.Conn(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if con != nil {
			con.Close()
		}
	}()
	db, err := con.ExecContext(ctx, insertScript, message)
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = db.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return err
	}

	return err
}

func (h *PostgresqlHandler) Handle(ctx context.Context, record slog.Record) error {

	message.Mu.Lock()
	defer message.Mu.Unlock()

	// Setting  user id & Trace Id here
	uId, tId := GetRquestAttributeFromContext(ctx)
	record.AddAttrs(slog.Attr{Key: "UserId", Value: slog.StringValue(uId)})
	record.AddAttrs(slog.Attr{Key: "TraceId", Value: slog.StringValue(tId)})

	err := h.JSONHandler.Handle(ctx, record)
	if err != nil {
		return err
	}
	return h.WriteToDb(message.Message, ctx)
}

func NewPostgresqlHandler(ctx context.Context, options HandlerOptions) *PostgresqlHandler {

	h := &PostgresqlHandler{
		CustomHandler: CustomHandler{
			HandlerOptions: options,
		},
		Ctx:         ctx,
		JSONHandler: slog.NewJSONHandler(&message, &options.HandlerOptions),
	}
	InitHandler(options.ConnectionString)
	return h
}

func InitHandler(ConnectionString string) {

	dbInitialize.Do(func() {
		con, err := sql.Open("postgres", ConnectionString)
		db = con
		if err != nil {
			panic(err)
		}
		createLogTableIfNotExists(db)
	})

}

func (h *PostgresqlHandler) WithAttrs(attrs []slog.Attr) slog.Handler {

	jsonHandler := h.JSONHandler.WithAttrs(attrs)
	newh := NewPostgresqlHandler(h.Ctx, h.HandlerOptions)

	newh.JSONHandler = jsonHandler.(*slog.JSONHandler)
	return newh
}

func (h *PostgresqlHandler) WithGroup(name string) slog.Handler {

	jsonHandler := h.JSONHandler.WithGroup(name)
	newh := NewPostgresqlHandler(h.Ctx, h.HandlerOptions)

	newh.JSONHandler = jsonHandler.(*slog.JSONHandler)
	return newh
}

func createLogTableIfNotExists(db *sql.DB) {

	createTableQuery := `CREATE SCHEMA IF NOT EXISTS log;
	CREATE TABLE IF NOT EXISTS log.log
	(
		ID  SERIAL PRIMARY KEY,
		message jsonb
	);
	`

	_, err := db.Exec(createTableQuery)

	if err != nil {
		panic(err)
	}

}

const insertScript string = `INSERT INTO "log"."log"( message)
	VALUES ($1);`
