package events

import (
	"context"
	"strings"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/google/uuid"

	//"github.com/neel4os/warg/internal/account-management/domain/account/app/commands_events"
	"github.com/neel4os/warg/internal/account-management/domain/account/app/events"
	"github.com/neel4os/warg/internal/account-management/domain/organization/app/commands"
	"github.com/neel4os/warg/internal/eventstore/domain/app"
	"github.com/rs/zerolog/log"
)

// CreateOrgOnAccountCreated is an event handler that handled AccountOnboarded event
// this just publish CreateOrgCommand to the command bus
type CreateOrgOnAccountCreatedEventHandler struct {
	commandBus *cqrs.CommandBus
}

func NewCreateOrgOnAccountCreatedEventHandler() *CreateOrgOnAccountCreatedEventHandler {
	ep := app.GetEventPlatform()
	return &CreateOrgOnAccountCreatedEventHandler{
		commandBus: ep.CommandBus,
	}
}

func (c *CreateOrgOnAccountCreatedEventHandler) Handle(ctx context.Context, event *events.AccountOnboarded) error {
	log.Info().Caller().Interface("Handling event CreateOrgOnAccountCreated ", &event).Msg("")
	org_id := uuid.New()
	// create org command
	email := event.Email
	domain := ""
	if atIdx := strings.LastIndex(email, "@"); atIdx != -1 {
		domain = email[atIdx+1:]
	}
	createOrgCommand := &commands.CreateOrgCommand{
		AccountId:  event.AccountId,
		OrgId:      org_id,
		OrgName:    event.AccountName,
		DomainName: domain,
	}
	// send command to command bus
	return c.commandBus.Send(ctx, createOrgCommand)
}
