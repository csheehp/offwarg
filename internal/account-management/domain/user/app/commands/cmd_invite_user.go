package commands

import (
	"context"
	"encoding/json"

	"github.com/neel4os/warg/internal/account-management/domain/user/aggregates/value"
	user_repo "github.com/neel4os/warg/internal/account-management/domain/user/repositories"
	persistence_user "github.com/neel4os/warg/internal/account-management/persistence/users"
	"github.com/neel4os/warg/internal/common/config"
	"github.com/neel4os/warg/internal/common/database"
	"github.com/neel4os/warg/internal/common/errors"
	"github.com/neel4os/warg/internal/common/smtp"
	"github.com/neel4os/warg/internal/eventstore/domain/aggregates"
	"github.com/neel4os/warg/internal/eventstore/domain/repositories"
	"github.com/neel4os/warg/internal/eventstore/persistence"
	"github.com/rs/zerolog/log"
	"gorm.io/datatypes"
)

type InviteUserCommand struct {
	Email  string
	UserId string
}

type InviteUserCommandHandler struct {
	eventRepo repositories.EventRepositories
	userRepo  user_repo.UserRepositoryInterface
}

func NewInviteUserCommandHandler() *InviteUserCommandHandler {
	_config := config.GetConfig()
	dbcon := database.GetDataConn(*_config)
	eventRepo := persistence.NewEventDatabaseRepository(dbcon)
	userRepo := persistence_user.NewUserConcreteRepository()
	return &InviteUserCommandHandler{
		eventRepo: eventRepo,
		userRepo:  userRepo,
	}
}

func (c *InviteUserCommandHandler) Handle(ctx context.Context, cmd *InviteUserCommand) error {
	log.Debug().Caller().Interface("Handling command invite user", &cmd).Msg("")
	userStream := value.GetUserStream()
	_event := aggregates.NewEvent(
		userStream.StreamID(),
		userStream.StreamName()+"."+cmd.UserId,
	)
	_req_bytes, err := json.Marshal(cmd)
	if err != nil {
		log.Error().Err(err).Caller().Msg("Error while marshalling event data")
		return errors.NewJSONMarhsalError(err.Error())
	}
	_event = _event.SetEventType("user_invited").
		SetEventData(datatypes.JSON(_req_bytes)).
		SetMetadata(datatypes.JSON{}).
		SetInitiatorType("system").
		SetInitiatorName("InviteUserCommandHandler")
	tx, err := c.eventRepo.CreateEvent(_event)
	if err != nil {
		log.Error().Err(err).Caller().Msg("Error while creating event")
		tx.Rollback()
		return errors.NewDatabaseOperationError(err.Error())
	}
	verification_cache, err := c.userRepo.CreateUserVerficationCache(cmd.UserId)
	if err != nil {
		log.Error().Err(err).Caller().Msg("Error while creating user verification cache")
		tx.Rollback()
		return errors.NewDatabaseOperationError(err.Error())
	}
	// send the mail
	_mail_client := smtp.NewSendMail()
	_mail_body := verification_cache
	err = _mail_client.SendMail(cmd.Email, "Invitation to join Warg", _mail_body)
	if err != nil {
		log.Error().Err(err).Caller().Msg("Error while sending mail")
		tx.Rollback()
		// FIXME: new error type
		return err
	}
	tx.Commit()
	return nil
}
