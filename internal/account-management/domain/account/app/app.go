package app

import "github.com/neel4os/warg/internal/account-management/domain/account/app/commands_events"

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	AccountOnboardCommand *commands_events.AccountOnboardingCommandHandler
}
type Queries struct{}
