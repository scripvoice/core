package infrastructure

import (
	"context"
	"fmt"
	"sync"
)

// Define the domain query interface
type DomainQuery interface {
	QueryName() string
}

// Define the domain query handler interface
type DomainQueryHandler interface {
	Execute(query DomainQuery, context context.Context) (interface{}, error)
}

// Define the domain query handler factory type
type DomainQueryHandlerFactory struct {
	handlers map[string]DomainQueryHandler
}

func NewDomainQueryHandlerFactory() *DomainQueryHandlerFactory {
	return &DomainQueryHandlerFactory{
		handlers: make(map[string]DomainQueryHandler),
	}
}

var (
	queryfactoryinstance *DomainQueryHandlerFactory
	oncequeryfactory     sync.Once
)

// GetInstance returns the singleton instance
func GetDomainQueryHandlerFactoryInstance() *DomainQueryHandlerFactory {
	oncequeryfactory.Do(func() {
		queryfactoryinstance = NewDomainQueryHandlerFactory() // Create the singleton instance
	})
	return queryfactoryinstance
}

// Register a new query handler with the factory
func (f *DomainQueryHandlerFactory) RegisterHandler(queryName string, handler DomainQueryHandler) {
	if f.handlers == nil {
		f.handlers = make(map[string]DomainQueryHandler)
	}
	f.handlers[queryName] = handler
}

// Resolve a query handler from the factory by its name
func (f *DomainQueryHandlerFactory) ResolveHandler(queryName string) (DomainQueryHandler, error) {
	handler, ok := f.handlers[queryName]
	if !ok {
		return nil, fmt.Errorf("no handler registered for query %q", queryName)
	}
	return handler, nil
}

type QueryRegistrator interface {
	Register(factory *DomainQueryHandlerFactory)
}
