package commands

import (
	"context"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/google/uuid"
	//"github.com/neel4os/warg/internal/account-management/domain/account/app/events"
	"github.com/neel4os/warg/internal/common/config"
	"github.com/neel4os/warg/internal/common/database"
	"github.com/neel4os/warg/internal/eventstore/domain/app"
	"github.com/neel4os/warg/internal/eventstore/domain/repositories"
	"github.com/neel4os/warg/internal/eventstore/persistence"
	"github.com/rs/zerolog/log"
)

type CreateOrgCommand struct {
	AccountId uuid.UUID `json:"account_id" valid:"uuid,required~account_id required and must be a valid uuid"`
	OrgId     uuid.UUID
	OrgName   string `json:"org_name" valid:"alphanum,required~org_name required and must be alphanumeric"`
	DomainName string 
}

type CreateOrgCommandHandler struct {
	eventRepo repositories.EventRepositories
	eventBus  *cqrs.EventBus
}

func NewCreateOrgCommandHandler() *CreateOrgCommandHandler {
	_config := config.GetConfig()
	dbcon := database.GetDataConn(*_config)
	eventRepo := persistence.NewEventDatabaseRepository(dbcon)
	eventPlatform := app.GetEventPlatform()
	return &CreateOrgCommandHandler{
		eventRepo: eventRepo,
		eventBus:  eventPlatform.EventBus,
	}
}

func (c *CreateOrgCommandHandler) Handle(ctx context.Context, cmd *CreateOrgCommand) error {
	log.Debug().Caller().Interface("Handling command org", &cmd).Msg("")
	
	return nil
}
