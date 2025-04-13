package aggregates

import (
	//"time"

	"github.com/google/uuid"
	"github.com/neel4os/warg/internal/account-management/domain/user/aggregates/value"
)

type User struct {
	ID        uuid.UUID        `json:"id"`
	FirstName string           `json:"first_name"`
	LastName  string           `json:"last_name"`
	Email     string           `json:"email"`
	Status    value.UserStatus `json:"status"`
	// CreatedAt time.Time        `json:"created_at"`
	// UpdatedAt time.Time        `json:"updated_at"`
	// IsManaged bool             `json:"is_managed"`
}

type UserCreated struct {
	AccountId uuid.UUID `json:"account_id"`
	OrgId     uuid.UUID `json:"org_id"`
	UserId    uuid.UUID `json:"user_id"`
	Email     string    `json:"email"`
}
