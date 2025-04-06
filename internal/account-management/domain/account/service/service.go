package service

import (
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	// "github.com/neel4os/warg/internal/account-management/domain/account/app"
	"github.com/neel4os/warg/internal/account-management/domain/account/app/commands_events"
	eventstore "github.com/neel4os/warg/internal/eventstore/domain/app"
	"github.com/rs/zerolog/log"
)

// func NewAccountApplication() *app.Application {
// 	return &app.Application{
// 		Commands: app.Commands{
// 			AccountOnboardCommand: commands_events.NewAccountOnboardCommandHandler(),
// 		},
// 		Queries: app.Queries{},
// 	}
// }

func RegisterCommandHandlers(ep *eventstore.EventPlatform) {
	log.Info().Str("component", "service").Msg("Registering command handlers")
	err := ep.AddCommandProcessorHandler(cqrs.NewCommandHandler("AccountOnboarded", commands_events.NewAccountOnboardCommandHandler().Handle))
	if err != nil {
		panic(err)
	}
}

func RegisterEventHandlers(ep *eventstore.EventPlatform) {
	log.Info().Str("component", "service").Msg("Registering event handlers")
	err := ep.AddEventProcessorHandler(cqrs.NewEventHandler("CreateOrgOnAccountCreated", commands_events.NewCreateOrgOnAccountCreatedEventHandler().Handle))
	if err != nil {
		panic(err)
	}
}
