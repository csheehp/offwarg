package eventstore

import "time"

type EventStore struct {
	ID         uint      `gorm:"primaryKey"`          // Auto-incrementing ID
	StreamID   string    `gorm:"not null"`            // Stream ID (e.g., aggregate ID)
	StreamName string    `gorm:"not null"`            // Stream name (e.g., aggregate type)
	EventType  string    `gorm:"not null"`            // Event type (e.g., "OrderCreated")
	EventData  string    `gorm:"type:jsonb;not null"` // Event payload (JSON)
	Metadata   string    `gorm:"type:jsonb"`          // Optional metadata (JSON)
	Version    int       `gorm:"not null"`            // Version of the stream
	CreatedAt  time.Time `gorm:"autoCreateTime"`
}
