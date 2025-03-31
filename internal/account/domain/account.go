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
	UpdatedAt time.Time     `json:"updated_at"`
	Name      string        `json:"name"`
	Status    AccountStatus `json:"status"`
}

func NewAccount(name, firstName, lastName, email string) *Account {
	return &Account{
		Name:   name,
		Status: AccountStatusPending,
	}
}

type AccountCreationRequest struct {
	AccountName string `json:"account_name" valid:"alphanum,required~account_name required and must be alphanumeric"`
	FirstName   string `json:"first_name" valid:"alpha,required~first_name required and must be alphabetic"`
	LastName    string `json:"last_name" valid:"alpha,required~last_name required and must be alphabetic"`
	Email       string `json:"email" valid:"email,required~email required and must be a valid email address"`
}
