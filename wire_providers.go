package core

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	)

func ProvideMySqlContext() *MySqlContext {
	return NewMySqlContext(viper.GetString("db_conn"))
	}

func ProvideJwtAuth() *JwtAuth {
	return NewJwtAuth()
}

	var DependencySet = wire.NewSet(
		ProvideMySqlContext, ProvideJwtAuth	)