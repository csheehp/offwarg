package domain

import (
	"time"

	"github.com/google/uuid"
)

type AccountStatus string

const (
	AccountStatusActive   AccountStatus = "active"
	AccountStatusInactive AccountStatus = "inactive"
	AccountStatusPending  AccountStatus = "created"
	AccountStatusDeleted  AccountStatus = "deleted"
)

func (s AccountStatus) IsValid() bool {
	switch s {
	case AccountStatusActive, AccountStatusInactive, AccountStatusPending, AccountStatusDeleted:
		return true
	}
	return false
}

func (s AccountStatus) String() string {
	switch s {
	case AccountStatusActive:
		return "active"
	case AccountStatusInactive:
		return "inactive"
	case AccountStatusPending:
		return "pending"
	case AccountStatusDeleted:
		return "deleted"
	default:
		return "unknown"
	}
}

// Account represents an account in the system.

type Account struct {
	ID        uuid.UUID     `json:"id"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	Name      string        `json:"name"`
	FisrtName string        `json:"first_name"`
	LastName  string        `json:"last_name"`
	Email     string        `json:"email"`
	Status    AccountStatus `json:"status"`
}

func NewAccount(name, firstName, lastName, email string) *Account {
	return &Account{
		Name:      name,
		FisrtName: firstName,
		LastName:  lastName,
		Email:     email,
		Status:    AccountStatusPending,
	}
}
