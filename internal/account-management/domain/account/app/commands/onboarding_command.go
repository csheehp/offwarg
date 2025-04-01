package commands

import (
	"encoding/json"

	"github.com/neel4os/warg/internal/account-management/domain/account/aggregates/value"
	account_repository "github.com/neel4os/warg/internal/account-management/domain/account/repositories"
	account_persistence "github.com/neel4os/warg/internal/account-management/persistence/account"
	"github.com/neel4os/warg/internal/common/errors"
	"github.com/neel4os/warg/internal/eventstore/domain/aggregates"
	domain_repository "github.com/neel4os/warg/internal/eventstore/domain/repositories"
	event_persistence "github.com/neel4os/warg/internal/eventstore/persistence"
	"gorm.io/datatypes"
)

// type AccountOnboardHandler decorators.CommandHandler[value.AccountCreationRequest]

type AccountOnboardingCommand struct {
	accountRepo account_repository.AccountRepositoryInterface
	eventRepo   domain_repository.EventRepositories
}

func NewAccountOnboardCommand() *AccountOnboardingCommand {
	accountRepo := account_persistence.NewAccountDatabaseRepository()
	eventRepo := event_persistence.NewEventDatabaseRepository()
	return &AccountOnboardingCommand{
		accountRepo: accountRepo,
		eventRepo:   eventRepo,
	}
}

func (h *AccountOnboardingCommand) Handle(req value.AccountCreationRequest) error {
	// get stream info of account stream
	accountStream := value.GetAccountStream()
	// create event
	_event := aggregates.NewEvent(
		accountStream.StreamID(),
		accountStream.StreamName(),
	)
	_req_bytes, err := json.Marshal(req)
	if err != nil {
		return errors.NewJSONMarhsalError(err.Error())
	}
	_event = _event.SetInitiatorType("user").
		SetInitiatorName(req.Email).
		SetEventData(datatypes.JSON{}).
		SetEventData(datatypes.JSON(_req_bytes))
	err = h.eventRepo.CreateEvent(_event)
	if err != nil {
		return err
	}
	return nil
}
