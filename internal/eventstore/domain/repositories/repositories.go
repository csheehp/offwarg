package repositories

import (
	"github.com/neel4os/warg/internal/eventstore/domain/aggregates"
)

type EventRepositories interface {
	CreateEvent(*aggregates.Event) error
}
