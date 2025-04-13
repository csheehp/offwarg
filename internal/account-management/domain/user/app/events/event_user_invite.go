package events

import (
	"context"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	organization "github.com/neel4os/warg/internal/account-management/domain/organization/aggregates"
	"github.com/neel4os/warg/internal/account-management/domain/user/app/commands"
	"github.com/neel4os/warg/internal/eventstore/domain/app"
	"github.com/rs/zerolog/log"
)

type InviteUserEventHandler struct {
	commandBus *cqrs.CommandBus
}

func NewInviteUserEventHandler() *InviteUserEventHandler {
	ep := app.GetEventPlatform()
	return &InviteUserEventHandler{
		commandBus: ep.CommandBus,
	}
}

func (c *InviteUserEventHandler) Handle(ctx context.Context, event *organization.OrganizationUpdatedUserMoved) error {
	log.Debug().Caller().Interface("Handling event InviteUserEventHandler ", &event).Msg("")
	inviteUserCmd := &commands.InviteUserCommand{
		Email:  event.Email,
		UserId: event.UserId.String(),
	}
	return c.commandBus.Send(ctx, inviteUserCmd)
}
