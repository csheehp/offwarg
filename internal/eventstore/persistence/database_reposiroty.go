package persistence

import "github.com/neel4os/warg/internal/eventstore/domain/aggregates"

type eventDatabaseRepository struct{}

func NewEventDatabaseRepository() *eventDatabaseRepository {
	return &eventDatabaseRepository{}
}

func (r *eventDatabaseRepository) CreateEvent(event *aggregates.Event) error {
	// Implement the logic to create an event in the database
	return nil
}
