package infrastructure

type IMediator interface {
	NotifyHandlers(domainEvents []DomainEvent)
}

type DomainEventMediator struct {
	eventFactory *EventFactory
}

// NewDomainEventMediator creates a new instance of DomainEventMediator.
func NewDomainEventMediator(eventFactory *EventFactory) *DomainEventMediator {
	return &DomainEventMediator{
		eventFactory: eventFactory,
	}
}

// NotifyHandlers notifies the event handlers based on the domain events received.
func (mediator *DomainEventMediator) NotifyHandlers(domainEvents []DomainEvent) {
	for _, event := range domainEvents {
		handler := mediator.eventFactory.ResolveEventHandler(event.GetName())
		if handler != nil {
			handler.HandleEvent(event)
		}
	}
}
