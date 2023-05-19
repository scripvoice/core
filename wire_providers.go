package core

import (
	"context"

	"github.com/google/wire"
	config "github.com/scripvoice/core/config"
	store "github.com/scripvoice/core/datastore"
	infra "github.com/scripvoice/core/infrastructure"
	"github.com/spf13/viper"
)

func ProvideMySqlContext() *MySqlContext {
	return NewMySqlContext(viper.GetString("ConnectionString"))
}

func ProvideDbContext(ctx context.Context) (*store.SqlStore, error) {
	return store.NewMysqlStore(ctx, config.Values.ConnectionString)
}

func ProvideJwtAuth() *JwtAuth {
	return NewJwtAuth()
}

func ProvideEventFactory() *infra.EventFactory {
	return infra.GetEventFactoryInstance()()
}

func ProvideDomainEventMediator(eventFactory *infra.EventFactory) *infra.DomainEventMediator {
	return infra.NewDomainEventMediator(eventFactory)
}

var DependencySet = wire.NewSet(
	ProvideMySqlContext, ProvideJwtAuth, ProvideDbContext, ProvideEventFactory, ProvideDomainEventMediator)
