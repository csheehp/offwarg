package start

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/neel4os/warg/internal/config"
	"github.com/neel4os/warg/internal/controller"
	"github.com/rs/zerolog/log"
)

type starter struct{
	cfg *config.Config
}

func NewStarter(cfg *config.Config) *starter {
	return &starter{cfg: cfg}
}

func (s *starter) DoStart() error {
	ctrlr := controller.NewController(s.cfg)
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGABRT)
	ctrlr.Init()
	ctrlr.Run()
	v := <-exit
	log.Info().Str("signal", v.String()).Caller().Msg("Received signal")
	ctrlr.Stop()
	return nil
}
