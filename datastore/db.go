package datastore

import (
	"context"
	"database/sql"
	"errors"

	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetMysqlDb(ConnectionString string) (*gorm.DB, error) {

	database, err := gorm.Open(mysql.Open(ConnectionString), &gorm.Config{
		SkipDefaultTransaction: true})
	if err != nil {
		return nil, err
	}
	return database, nil
}

type SqlStore struct {
	db             *gorm.DB
	HasTransaction bool
}

func (s *SqlStore) DB() *gorm.DB {
	return s.db
}

func NewMysqlStore(ctx context.Context, ConnectionString string) (*SqlStore, error) {
	db, err := GetMysqlDb(ConnectionString)
	if err != nil {
		return nil, err
	}
	return &SqlStore{db: db.WithContext(ctx)}, nil
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

func GetRepo[T Irepo](s *SqlStore) (T, error) {
	t := new(T)
	var r Irepo = (*t).SetStore(s)
	v, ok := r.(T)
	if !ok {
		return *t, errors.New("Unable to get base repo")
	}

	return v, nil
}

type Irepo interface {
	SetStore(store *SqlStore) Irepo
}

type BaseRepo[T any] struct {
	Store     *SqlStore
	storeInit sync.Once
}

func (r BaseRepo[T]) SetStore(s *SqlStore) Irepo {
	r.storeInit.Do(func() {
		r.Store = s
	})
	return r
}

func (r BaseRepo[T]) Save(entity T) {
	r.Store.db.Save(entity)
}
