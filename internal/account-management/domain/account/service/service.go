package service

import (
	"github.com/neel4os/warg/internal/account-management/domain/account/app"
	"github.com/neel4os/warg/internal/account-management/domain/account/app/commands"
)

func NewAccountApplication() *app.Application {
	return &app.Application{
		Commands: app.Commands{
			AccountOnboardCommand: commands.NewAccountOnboardCommand(),
		},
		Queries: app.Queries{},
	}
}
