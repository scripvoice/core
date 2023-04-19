package core

import (
	"context"
	"github.com/google/wire"
	"github.com/spf13/viper"
	store "github.com/core/infrastructure/datastore/mysql"
	config "github.com/core/config"
	)

func ProvideMySqlContext() *MySqlContext {
	return NewMySqlContext(viper.GetString("db_conn"))
	}

func ProvideDbContext(ctx context.Context) (*store.SqlStore, error)  {
	return NewMysqlStore(config.Values.ConnectionString)
}

func ProvideJwtAuth() *JwtAuth {
	return NewJwtAuth()
}

	var DependencySet = wire.NewSet(
		ProvideMySqlContext, ProvideJwtAuth, ProvideDbContext	)