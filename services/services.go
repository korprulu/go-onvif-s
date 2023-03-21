// Package services ...
package services

import (
	"context"
	"net/http"

	"github.com/jfsmig/onvif/networking"
	"github.com/korprulu/go-onvif-s/services/device"
	"github.com/korprulu/go-onvif-s/services/media"
)

// Services ...
type Services struct {
	Device device.API
	Media  media.API

	rtpOverTCP   bool
	rtpMulticast bool
	rtsp         bool
}

// Config ...
type Config struct {
	IPAddr     string
	UUID       string
	Username   string
	Password   string
	HTTPClient *http.Client
}

// New creates a new Services instance
func New(ctx context.Context, cfg Config) (*Services, error) {
	// the device endpoint will be added when create new client
	client, err := networking.NewClient(networking.ClientInfo{
		Xaddr: cfg.IPAddr,
		Uuid:  cfg.UUID,
	}, cfg.HTTPClient)
	if err != nil {
		return nil, err
	}

	client.SetAuth(networking.ClientAuth{
		Username: cfg.Username,
		Password: cfg.Password,
	})

	dev, err := device.New(ctx, client)
	if err != nil {
		return nil, err
	}

	capabilities, err := dev.GetCapabilities(ctx, device.GetCapabilitiesInput{Category: "All"})
	if err != nil {
		return nil, err
	}

	media := loadMedia(capabilities, client)

	return &Services{
		Device:       dev,
		Media:        media,
		rtpOverTCP:   bool(capabilities.Media.StreamingCapabilities.RTP_TCP),
		rtpMulticast: bool(capabilities.Media.StreamingCapabilities.RTPMulticast),
		rtsp:         bool(capabilities.Media.StreamingCapabilities.RTP_RTSP_TCP),
	}, nil
}

func loadMedia(capabilities *device.Capabilities, client *networking.Client) media.API {
	endpoint := string(capabilities.Media.XAddr)
	if len(endpoint) == 0 {
		return nil
	}
	client.AddEndpoint("media", endpoint)
	return media.New(client)
}

// RTPOverTCP ...
func (svc *Services) RTPOverTCP() bool {
	return svc.rtpOverTCP
}

// RTPMulticast ...
func (svc *Services) RTPMulticast() bool {
	return svc.rtpMulticast
}

// RTSP ...
func (svc *Services) RTSP() bool {
	return svc.rtsp
}
