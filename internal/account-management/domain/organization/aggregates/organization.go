package organization

import (
	"time"

	"github.com/neel4os/warg/internal/account-management/domain/user/aggregates"
)

type Organization struct {
	ID        string          `json:"id"`
	Name      string          `json:"name"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	Status    string          `json:"status"`
	Owner     aggregates.User `json:"owner"`
}
