package commands

import (
	"context"
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	//"github.com/neel4os/warg/internal/account-management/domain/account/aggregates/value"
	//"github.com/neel4os/warg/internal/account-management/domain/account/aggregates"
	"github.com/neel4os/warg/internal/account-management/domain/account/aggregates/value"
	//"github.com/neel4os/warg/internal/account-management/domain/account/app/events"
	account_repository "github.com/neel4os/warg/internal/account-management/domain/account/repositories"
	account_persistence "github.com/neel4os/warg/internal/account-management/persistence/account"

	//"github.com/neel4os/warg/internal/common/database"
	"github.com/neel4os/warg/internal/common/errors"
	"github.com/neel4os/warg/internal/eventstore/domain/aggregates"
	"github.com/neel4os/warg/internal/eventstore/domain/app"
	domain_repository "github.com/neel4os/warg/internal/eventstore/domain/repositories"
	event_persistence "github.com/neel4os/warg/internal/eventstore/persistence"
	"gorm.io/datatypes"
)

// type AccountOnboardHandler decorators.CommandHandler[value.AccountCreationRequest]

type OnBoardAccount struct {
	AccountName string `json:"account_name" valid:"alphanum,required~account_name required and must be alphanumeric"`
	FirstName   string `json:"first_name" valid:"alpha,required~first_name required and must be alphabetic"`
	LastName    string `json:"last_name" valid:"alpha,required~last_name required and must be alphabetic"`
	Email       string `json:"email" valid:"email,required~email required and must be a valid email address"`
}

type AccountOnboardingCommandHandler struct {
	accountRepo account_repository.AccountRepositoryInterface
	eventRepo   domain_repository.EventRepositories
	eventBus    *cqrs.EventBus
}

func NewAccountOnboardCommandHandler() *AccountOnboardingCommandHandler {
	accountRepo := account_persistence.NewAccountDatabaseRepository()
	eventRepo := event_persistence.NewEventDatabaseRepository()
	eventPlatform := app.GetEventPlatform()
	return &AccountOnboardingCommandHandler{
		accountRepo: accountRepo,
		eventRepo:   eventRepo,
		eventBus:    eventPlatform.EventBus,
	}
}

func (h *AccountOnboardingCommandHandler) Handle(ctx context.Context, cmd *OnBoardAccount) error {
	log.Debug().Caller().Interface("Handling command", &cmd).Msg("")
	// return h.eventBus.Publish(ctx, &events.AccountOnboarded{})
	// create account ID
	_account_ID := uuid.New()
	// get stream info of account stream
	accountStream := value.GetAccountStream()
	// create event
	_event := aggregates.NewEvent(
		accountStream.StreamID(),
		accountStream.StreamName() + "." + _account_ID.String(),
	)
	_req_bytes, err := json.Marshal(cmd)
	if err != nil {
		return errors.NewJSONMarhsalError(err.Error())
	}
	// make event more informative
	_event = _event.SetInitiatorType("user").
		SetInitiatorName(cmd.Email).
		SetMetadata(datatypes.JSON{}).
		SetEventData(datatypes.JSON(_req_bytes)).
		SetEventType("account_onboarded")
	// Now lets get the database connection
	//dbcon := database.GetDataConn()
	// err = h.eventRepo.CreateEvent(_event)
	// if err != nil {
	// 	return err
	// }
	return nil
}
