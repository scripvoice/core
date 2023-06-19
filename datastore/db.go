package datastore

import (
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetMysqlDb(ConnectionString string, c gorm.Config) (*gorm.DB, error) {

	database, err := gorm.Open(mysql.Open(ConnectionString), &c)
	if err != nil {
		return nil, err
	}
	return database, nil
}

func GetPostgreSqlDb(ConnectionString string, c gorm.Config) (*gorm.DB, error) {

	database, err := gorm.Open(postgres.Open(ConnectionString), &c)
	if err != nil {
		return nil, err
	}
	return database, nil
}
