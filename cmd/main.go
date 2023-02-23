package main

import (
	"github.com/korprulu/go-onvif-s/internal/discovery"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	prober := &discovery.Prober{}

	devices, err := prober.Probe()
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}

	for _, device := range devices {
		log.Info().Msgf("devices: %#v", *device)
	}
}
