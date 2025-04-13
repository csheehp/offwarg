package initialize

import (
	// "errors"
	// "time"

	// "github.com/neel4os/warg/internal/common/cache"
	"github.com/neel4os/warg/internal/common/config"
	"github.com/neel4os/warg/internal/common/util"
	"github.com/rs/zerolog/log"
)

type idpVerifier struct {
	idpConfig config.IdpConfig
	realmUrl  string
}

func newidpVerifier(cfg config.IdpConfig) *idpVerifier {
	return &idpVerifier{idpConfig: cfg, realmUrl: cfg.Url + "/realms/" + cfg.RealmName}
}

func (i *idpVerifier) Verify() error {
	// verify if we can access well-known configuration
	// if we can't, return error and true
	client := util.NewRestClient()
	log.Debug().Caller().Msgf("Checking well known %s", i.idpConfig.IdpName)
	_, err := client.R().Get(i.realmUrl + "/.well-known/openid-configuration")
	if err != nil {
		log.Error().Err(err).Msgf("Failed to get well-known configuration for %s", i.idpConfig.IdpName)
		return err
	}
	// log.Debug().Caller().Msg("Checking token updating....")
	// cfg := config.GetConfig()
	// cache := cache.NewIMCache(cfg)
	// token := cache.GetToken()
	// _token := ""
	// if token == "" {
	// 	log.Error().Msgf("Failed to get token for %s", i.idpConfig.IdpName)
	// 	return errors.New("Failed to get token for " + i.idpConfig.IdpName)
	// }
	// for range 12 {
	// 	_token = cache.GetToken()
	// 	log.Info().Msgf("Token for %s is %s", i.idpConfig.IdpName, _token)
	// 	time.Sleep(5 * time.Second)
	// }
	// if token == _token {
	// 	log.Error().Msgf("Failed to update token for %s", i.idpConfig.IdpName)
	// 	return errors.New("Failed to update token for " + i.idpConfig.IdpName)
	// }

	return nil
}

func (i *idpVerifier) Name() string {
	return i.idpConfig.IdpName
}
