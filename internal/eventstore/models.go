package eventstore

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Event struct {
	ID            uuid.UUID      `gorm:"primaryKey"`          // Auto-incrementing ID
	StreamID      uuid.UUID      `gorm:"not null"`            // Stream ID (e.g., aggregate ID)
	StreamName    string         `gorm:"not null"`            // Stream name (e.g., aggregate type)
	EventType     string         `gorm:"not null"`            // Event type (e.g., "OrderCreated")
	EventData     datatypes.JSON `gorm:"type:jsonb;not null"` // Event payload (JSON)
	Metadata      datatypes.JSON `gorm:"type:jsonb"`          // Optional metadata (JSON)
	Version       int            `gorm:"not null"`
	InitiatorType string         `gorm:"not null"`
	InitiatorName string         `gorm:"not null"`
	CreatedAt     time.Time      `gorm:"autoCreateTime"`
}
