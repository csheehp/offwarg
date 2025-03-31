package commands

import (
	"github.com/neel4os/warg/internal/account-management/domain/account/aggregates/value"
	"github.com/neel4os/warg/internal/account-management/domain/account/repositories"
	account_persistence "github.com/neel4os/warg/internal/account-management/persistence/account"
)

// type AccountOnboardHandler decorators.CommandHandler[value.AccountCreationRequest]

type AccountOnboardingCommand struct {
	accountRepo repositories.AccountRepositoryInterface
}

func NewAccountOnboardCommand() *AccountOnboardingCommand {
	accountRepo := account_persistence.NewAccountDatabaseRepository()
	return &AccountOnboardingCommand{
		accountRepo: accountRepo,
	}
}

func (h *AccountOnboardingCommand) Handle(req value.AccountCreationRequest) error {
	err := h.accountRepo.CreateAccount(req)
	if err != nil {
		return err
	}
	return nil
}
