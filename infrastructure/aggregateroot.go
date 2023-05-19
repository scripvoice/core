package infrastructure

// AggregateRoot defines the base struct for aggregate roots.
type AggregateRoot struct {
	ID      int           `gorm:"-"`
	Changes []DomainEvent `gorm:"-"`
}

// ApplyChange applies a domain event to the aggregate root.
func (ar *AggregateRoot) ApplyChange(event DomainEvent) {
	if ar.Changes == nil {
		ar.Changes = make([]DomainEvent, 0)
	}
	ar.Changes = append(ar.Changes, event)
}

// RemoveDomainEvent removes a domain event from the aggregate root's changes.
func (ar *AggregateRoot) RemoveDomainEvent(event DomainEvent) {
	for i, e := range ar.Changes {
		if e.GetName() == event.GetName() {
			ar.Changes = append(ar.Changes[:i], ar.Changes[i+1:]...)
			break
		}
	}
}

// ClearDomainEvents clears all domain events from the aggregate root's changes.
func (ar *AggregateRoot) ClearDomainEvents() {
	ar.Changes = []DomainEvent{}
}
