package infrastructure

// AggregateRoot defines the base struct for aggregate roots.
type AggregateRoot struct {
	ID      int
	Changes []DomainEvent
}

// ApplyChange applies a domain event to the aggregate root.
func (ar *AggregateRoot) ApplyChange(event DomainEvent) {
	ar.Changes = append(ar.Changes, event)
}
