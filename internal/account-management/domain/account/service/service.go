package service

import (
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/neel4os/warg/internal/account-management/domain/account/app"
	"github.com/neel4os/warg/internal/account-management/domain/account/app/commands"
	eventstore "github.com/neel4os/warg/internal/eventstore/domain/app"
	"github.com/rs/zerolog/log"
)

func NewAccountApplication() *app.Application {
	return &app.Application{
		Commands: app.Commands{
			AccountOnboardCommand: commands.NewAccountOnboardCommandHandler(),
		},
		Queries: app.Queries{},
	}
}

func RegisterCommandHandlers(ep *eventstore.EventPlatform) {
	log.Info().Str("component", "service").Msg("Registering command handlers")
	err := ep.AddCommandProcessorHandler(cqrs.NewCommandHandler("AccountOnboarded", commands.NewAccountOnboardCommandHandler().Handle))
	if err != nil {
		panic(err)
	}
}
