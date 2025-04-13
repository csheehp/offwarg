package initialize

import (
	"errors"

	"github.com/neel4os/warg/internal/common/config"
	"github.com/rs/zerolog/log"
)

type verifier interface {
	Verify() error
	Name() string
}

type wargInitialization struct {
	verifiers []verifier
}

func NewInitilizer(cfg *config.Config) *wargInitialization {
	_verifiers := make([]verifier, 0)
	_verifiers = append(_verifiers, newidpVerifier(cfg.IdpConfig))
	return &wargInitialization{verifiers: _verifiers}
}

func (w *wargInitialization) DoInitialize() error {
	var isError bool
	for _, v := range w.verifiers {
		log.Info().Str("verifier", v.Name()).Caller().Msg("Verifying " + v.Name())
		err := v.Verify()
		if err != nil {
			log.Err(err).Caller().Msg("Verification Error")
			isError = true
		}
	}
	if isError {
		log.Error().Caller().Msg("Warg initialization failed")
		return errors.New("Warg initialization failed")
	}
	log.Info().Caller().Msg("Warg initialization completed successfully")
	return nil
}
