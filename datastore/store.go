package datastore

import (
	"context"
	"database/sql"

	"gorm.io/gorm"
)

type SqlStore struct {
	db             *gorm.DB
	HasTransaction bool
}

func (s *SqlStore) Db() *gorm.DB {
	return s.db
}

func (s *SqlStore) DbWithContext(ctx context.Context) *gorm.DB {
	return s.db.WithContext(ctx)
}

func NewSqlStore(c SqlDbconfig) (*SqlStore, error) {

	var db *gorm.DB
	var err error

	switch c.DbType {
	case Mysql:
		db, err = GetMysqlDb(c.ConnectionString, c.GromConfig)
	case Postgres:
		db, err = GetPostgreSqlDb(c.ConnectionString, c.GromConfig)
	default:
		db, err = GetPostgreSqlDb(c.ConnectionString, c.GromConfig)

	}
	if err != nil {
		return nil, err
	}
	return &SqlStore{db: db}, nil
}

func (s *SqlStore) BeginTran(opts *sql.TxOptions) *gorm.DB {
	if !s.HasTransaction {
		s.HasTransaction = true
		s.db = s.db.Begin(opts)
	}
	return s.db
}

func (s *SqlStore) Commit() *gorm.DB {
	s.HasTransaction = false
	return s.db.Commit()
}

func (s *SqlStore) Rollback() *gorm.DB {
	s.HasTransaction = false
	return s.db.Rollback()
}
