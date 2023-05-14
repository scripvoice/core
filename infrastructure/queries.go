package infrastructure

import "fmt"

// Define the domain query interface
type DomainQuery interface {
	QueryName() string
}

// Define the domain query handler interface
type DomainQueryHandler interface {
	Execute(query DomainQuery) (interface{}, error)
}

// Define the domain query handler factory type
type DomainQueryHandlerFactory struct {
	handlers map[string]DomainQueryHandler
}

// Register a new query handler with the factory
func (f *DomainQueryHandlerFactory) RegisterHandler(handler DomainQueryHandler) {
	if f.handlers == nil {
		f.handlers = make(map[string]DomainQueryHandler)
	}
	f.handlers[handler.QueryName()] = handler
}

// Resolve a query handler from the factory by its name
func (f *DomainQueryHandlerFactory) ResolveHandler(queryName string) (DomainQueryHandler, error) {
	handler, ok := f.handlers[queryName]
	if !ok {
		return nil, fmt.Errorf("no handler registered for query %q", queryName)
	}
	return handler, nil
}
