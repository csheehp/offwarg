package controller

import (
	"github.com/neel4os/warg/internal/common/config"
	"github.com/rs/zerolog/log"
)

type controller struct {
	components []componentable
	cfg *config.Config
}

func NewController(cfg *config.Config) *controller {
	_components := make([]componentable, 0)
	_components = append(_components, NewHTTPComponent(cfg))
	return &controller{components: _components, cfg: cfg}
}

func (c *controller) Init() {
	for _, comp := range c.components {
		log.Debug().Str("component", comp.Name()).Caller().Msg("Initializing " + comp.Name())
		comp.Init()
	}
}

func (c *controller) Run() {
	for _, comp := range c.components {
		log.Debug().Str("component", comp.Name()).Caller().Msg("Running " + comp.Name())
		go comp.Run()
	}
}

func (c *controller) Stop(){
	for _, comp := range c.components {
		log.Debug().Str("component", comp.Name()).Caller().Msg("Stopping " + comp.Name())
		comp.Stop()
	}
}
