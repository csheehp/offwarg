package commands_events

import (
	"context"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/google/uuid"
	"github.com/neel4os/warg/internal/eventstore/domain/app"
	"github.com/rs/zerolog/log"
)

// CreateOrgOnAccountCreated is an event handler that handled AccountOnboarded event
// and creates an organization for the account
type CreateOrgOnAccountCreatedEventHandler struct {
	commandBus *cqrs.CommandBus
}

func NewCreateOrgOnAccountCreatedEventHandler() *CreateOrgOnAccountCreatedEventHandler {
	ep := app.GetEventPlatform()
	return &CreateOrgOnAccountCreatedEventHandler{
		commandBus: ep.CommandBus,
	}
}

func (c *CreateOrgOnAccountCreatedEventHandler) Handle(ctx context.Context, event *AccountOnboarded) error {
	log.Info().Caller().Interface("Handling event CreateOrgOnAccountCreated ", &event).Msg("")
	org_id := uuid.New()
	// create org command
	createOrgCommand := &CreateOrgCommand{
		AccountId: event.AccountId,
		OrgId:     org_id,
		OrgName:   event.AccountName,
	}
	// send command to command bus
	return c.commandBus.Send(ctx, createOrgCommand)
}
