package core

import (
	"github.com/google/wire"
	Auth "github.com/scripvoice/core/auth"
	infra "github.com/scripvoice/core/infrastructure"
)

func ProvideJwtAuth() *Auth.JwtAuth {
	return Auth.NewJwtAuth()
}

func ProvideEventFactory() *infra.EventFactory {
	return infra.GetEventFactoryInstance()
}

func ProvideCommandFactory() *infra.CommandFactory {
	return infra.GetCommandFactoryInstance()
}

func ProvideQueryFactory() *infra.DomainQueryHandlerFactory {
	return infra.GetDomainQueryHandlerFactoryInstance()
}

func ProvideDomainEventMediator(eventFactory *infra.EventFactory) *infra.DomainEventMediator {
	return infra.NewDomainEventMediator(eventFactory)
}

var DependencySet = wire.NewSet(ProvideJwtAuth, ProvideEventFactory,
	ProvideDomainEventMediator, ProvideCommandFactory, ProvideQueryFactory)
