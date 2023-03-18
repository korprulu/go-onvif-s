// Package main ...
package main

import (
	"context"
	"os"
	"time"

	"github.com/korprulu/go-onvif-s/services"
	"github.com/rs/zerolog"
)

var (
	logger = zerolog.
		New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
		With().Timestamp().
		Logger()
)

func main() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	ctx := context.Background()

	srv, err := services.New(ctx, services.DeviceInfo{Addr: "192.168.1.102"}, nil, services.WithAuth("admin", "590885"))
	if err != nil {
		logger.Fatal().Err(err).Send()
	}

	cfgs, err := srv.Media.GetVideoEncoderConfigurations(ctx)
	if err != nil {
		logger.Error().Err(err).Send()
	}

	for _, cfg := range cfgs {
		guaranteedNumber, err := srv.Media.GetGuaranteedNumberOfVideoEncoderInstances(ctx, cfg.Token)
		if err != nil {
			logger.Error().Err(err).Send()
			continue
		}
		logger.Info().Msgf("%#v\n", guaranteedNumber)
	}
}
