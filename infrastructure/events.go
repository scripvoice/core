package infrastructure

import (
	"sync"
)

// DomainEvent defines the interface for domain events.
type DomainEvent interface {
	GetName() string
}

// EventHandler defines the interface for event handlers.
type EventHandler interface {
	HandleEvent(event DomainEvent)
}

// EventFactory represents the factory class that resolves the event handler based on the event type.
type EventFactory struct {
	eventHandlers map[string]EventHandler
}

// NewEventFactory creates a new instance of the EventFactory.
func NewEventFactory() *EventFactory {
	return &EventFactory{
		eventHandlers: make(map[string]EventHandler),
	}
}

var (
	instance *EventFactory
	once     sync.Once
)

// GetInstance returns the singleton instance
func GetEventFactoryInstance() *EventFactory {
	once.Do(func() {
		instance = NewEventFactory() // Create the singleton instance
	})
	return instance
}

// RegisterEventHandler registers an event handler for a specific event type.
func (factory *EventFactory) RegisterEventHandler(eventType string, handler EventHandler) {
	factory.eventHandlers[eventType] = handler
}

// ResolveEventHandler resolves the event handler based on the event type.
func (factory *EventFactory) ResolveEventHandler(eventType string) EventHandler {
	handler, ok := factory.eventHandlers[eventType]
	if ok {
		return handler
	}
	return nil
}

type EventRegistrator interface {
	Register(factory *EventFactory)
}
