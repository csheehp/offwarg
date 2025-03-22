package config

import (
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/neel4os/warg/internal/logging"
	"github.com/rs/zerolog/log"
)

type IdpName string

const (
	keycloak IdpName = "keycloak"
)

func (i IdpName) String() string {
	return string(i)
}

type Config struct {
	IdpConfig    IdpConfig
	LoggerConfig LoggerConfig
	ServerConfig ServerConfig
}

type ServerConfig struct {
	Port                  string `env:"WARG_SERVERCONFIG_PORT" envDefault:"9999"`
	HidePortInStdOut      bool   `env:"WARG_SERVERCONFIG_HIDE_PORT_IN_STDOUT" envDefault:"true"`
	ReadTimeout           int    `envDefault:"10"`
	WriteTimeout          int    `envDefault:"10"`
	GraceFullShutdownTime int    `envDefault:"10"`
}

type LoggerConfig struct {
	IsDebugLog bool `env:"WARG_LOGGERCONFIG_IS_DEBUG_LOG" envDefault:"true"`
}

type IdpConfig struct {
	IdpName   IdpName `env:"WARG_IDPCONFIG_IDP_NAME" envDefault:"keycloak"`
	Url       string  `env:"WARG_IDPCONFIG_IDP_URL" envDefault:"http://localhost:8080"`
	RealmName string  `env:"WARG_IDPCONFIG_IDP_REALM_NAME" envDefault:"warg"`
}

func New() *Config {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		log.Error().Err(err).Caller().Msg("failed to parse environment variables")
		os.Exit(1)
	}
	if cfg.LoggerConfig.IsDebugLog {
		logging.SetLogConfig(true)
	} else {
		logging.SetLogConfig(false)
	}
	return &cfg
}
