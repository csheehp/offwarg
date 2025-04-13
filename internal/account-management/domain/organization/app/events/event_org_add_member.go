package events

import (
	"context"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/neel4os/warg/internal/account-management/domain/organization/app/commands"
	"github.com/neel4os/warg/internal/account-management/domain/user/aggregates"
	"github.com/neel4os/warg/internal/eventstore/domain/app"
	"github.com/rs/zerolog/log"
)

type MoveUserIntoOrg struct {
	commandBus *cqrs.CommandBus
}

func NewMoveUserIntoOrgEventHandler() *MoveUserIntoOrg {
	ep := app.GetEventPlatform()
	return &MoveUserIntoOrg{
		commandBus: ep.CommandBus,
	}
}

func (c *MoveUserIntoOrg) Handle(ctx context.Context, event *aggregates.UserCreated) error {
	log.Info().Caller().Interface("Handling event MoveUserIntoOrg ", &event).Msg("")
	moveUserToOrgCmd := &commands.MoveUserIntoOrgCommand{
		AccountId: event.AccountId,
		OrgId:     event.OrgId,
		UserId:    event.UserId,
		Email:     event.Email,
	}
	// send command to command bus
	return c.commandBus.Send(ctx, moveUserToOrgCmd)
}
