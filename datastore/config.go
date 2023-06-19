package datastore

import "gorm.io/gorm"

type SqlDb string

const (
	Mysql    SqlDb = "MySql"
	Postgres SqlDb = "Postgres"
)

type SqlDbconfig struct {
	DbType           SqlDb
	ConnectionString string
	GromConfig       gorm.Config
}
