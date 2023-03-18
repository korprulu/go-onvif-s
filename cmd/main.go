// Package main ...
package main

import (
	"context"

	"github.com/korprulu/go-onvif-s/services/device"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	ctx := context.Background()

	dev, err := device.New(device.Info{Addr: "192.168.1.102"}, nil, device.WithAuth("admin", "590885"))
	if err != nil {
		log.Fatal().Err(err).Send()
	}

}
