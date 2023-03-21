// Package main ...
package main

import (
	"context"
	"os"
	"time"

	"github.com/korprulu/go-onvif-s/client"
	"github.com/korprulu/go-onvif-s/services"
	"github.com/rs/zerolog"
)

var logger = zerolog.
	New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
	With().Timestamp().
	Logger()

func main() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	ctx := context.Background()

	srv, err := services.New(ctx, services.Config{
		IPAddr:   "192.168.1.102",
		Username: "admin",
		Password: "590885",
	})
	if err != nil {
		logger.Fatal().Err(err).Send()
	}

	client.New(ctx, srv)
}
