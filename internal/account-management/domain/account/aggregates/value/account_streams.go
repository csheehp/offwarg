package value

import (
	"sync"

	"github.com/google/uuid"
)

type accountStream struct {
	streamID   uuid.UUID
	streamName string
}

var (
	instance *accountStream
	once     sync.Once
)

func GetAccountStream() *accountStream {
	once.Do(func() {
		instance = &accountStream{
			streamID:   uuid.New(),
			streamName: "account",
		}
	})
	return instance
}

func (a *accountStream) StreamID() uuid.UUID {
	return a.streamID
}

func (a *accountStream) StreamName() string {
	return a.streamName
}
