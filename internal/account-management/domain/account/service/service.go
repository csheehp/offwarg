package service

import "github.com/neel4os/warg/internal/account-management/domain/account/app"

func NewAccountApplication() *app.Application {
	return &app.Application{
		Commands: app.Commands{},
		Queries:  app.Queries{},
	}
}
