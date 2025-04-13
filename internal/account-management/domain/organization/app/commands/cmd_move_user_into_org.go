package commands

import (
	"context"
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/google/uuid"
	organization "github.com/neel4os/warg/internal/account-management/domain/organization/aggregates"
	"github.com/neel4os/warg/internal/account-management/domain/organization/aggregates/value"
	persistence_organization "github.com/neel4os/warg/internal/account-management/persistence/organization"
	"github.com/neel4os/warg/internal/common/config"
	"github.com/neel4os/warg/internal/common/database"
	"github.com/neel4os/warg/internal/common/errors"
	"github.com/neel4os/warg/internal/eventstore/domain/aggregates"
	"github.com/neel4os/warg/internal/eventstore/domain/app"
	"github.com/neel4os/warg/internal/eventstore/domain/repositories"
	"github.com/neel4os/warg/internal/eventstore/persistence"
	"github.com/rs/zerolog/log"
	"gorm.io/datatypes"
)

type MoveUserIntoOrgCommand struct {
	AccountId uuid.UUID `json:"account_id" valid:"uuid,required~account_id required and must be a valid uuid"`
	OrgId     uuid.UUID `json:"org_id" valid:"uuid,required~org_id required and must be a valid uuid"`
	UserId    uuid.UUID `json:"user_id" valid:"uuid,required~user_id required and must be a valid uuid"`
	Email string `json:"email" valid:"email,required~email required and must be a valid email"`
}

type MoveUserIntoOrgCommandHandler struct {
	eventRepo repositories.EventRepositories
	eventBus  *cqrs.EventBus
	idpRepo   *persistence_organization.OrganizationKeycloakRepository
}

func NewMoveUserIntoOrgCommandHandler() *MoveUserIntoOrgCommandHandler {
	_config := config.GetConfig()
	dbcon := database.GetDataConn(*_config)
	eventRepo := persistence.NewEventDatabaseRepository(dbcon)
	eventPlatform := app.GetEventPlatform()
	idpRepo := persistence_organization.NewOrganizationKeycloakRepository()
	return &MoveUserIntoOrgCommandHandler{
		eventRepo: eventRepo,
		eventBus:  eventPlatform.EventBus,
		idpRepo:   idpRepo,
	}
}

func (c *MoveUserIntoOrgCommandHandler) Handle(ctx context.Context, cmd *MoveUserIntoOrgCommand) error {
	log.Debug().Caller().Interface("Handling command MoveUserIntoOrg", &cmd).Msg("")
	//FIXME: this is not stateless
	// check if the user is already in the org
	// do the actual movement
	err := c.idpRepo.AddMemberInOrganization(cmd.OrgId.String(), cmd.UserId.String())
	if err != nil {
		log.Error().Err(err).Caller().Msg("Error while moving user into organization")
		return err
	}
	// create the event
	orgStream := value.GetOrganizationStream()
	_event := aggregates.NewEvent(
		orgStream.StreamID(),
		orgStream.StreamName()+"."+cmd.OrgId.String(),
	)
	_event_data, err := json.Marshal(cmd)
	if err != nil {
		log.Error().Err(err).Caller().Msg("Error while marshalling event data")
		return errors.NewJSONMarhsalError(err.Error())
	}
	_event = _event.SetInitiatorType(string(aggregates.InitiatorTypeSystem)).
		SetInitiatorName("MoveUserIntoOrgCommandHandler").
		SetEventType("organization_updated").
		SetMetadata(datatypes.JSON{}).
		SetEventData(datatypes.JSON(_event_data))
	tx, err := c.eventRepo.CreateEvent(_event)
	if err != nil {
		log.Error().Err(err).Caller().Msg("Error while creating event")
		tx.Rollback()
		return errors.NewDatabaseOperationError(err.Error())
	}
	err = c.eventBus.Publish(ctx, organization.OrganizationUpdatedUserMoved{
		OrgId:  cmd.OrgId,
		UserId: cmd.UserId,
		Email: cmd.Email,
	})
	if err != nil {
		log.Error().Err(err).Caller().Msg("Error while publishing event")
		tx.Rollback()
		return errors.NewDatabaseOperationError(err.Error())
	}
	tx.Commit()
	return nil
}
