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
}

// Option ...
type Option func(*networking.Client) error

// DeviceInfo ...
type DeviceInfo struct {
	Addr string
	UUID string
}

// New creates a new Services instance
func New(ctx context.Context, info DeviceInfo, httpClient *http.Client, options ...Option) (*Services, error) {
	clientInfo := networking.ClientInfo{
		Xaddr: info.Addr,
		Uuid:  info.UUID,
	}

	// the device endpoint will be added when create new client
	client, err := networking.NewClient(clientInfo, httpClient)
	if err != nil {
		return nil, err
	}

	for _, opt := range options {
		if err := opt(client); err != nil {
			return nil, err
		}
	}

	device, err := device.New(ctx, client)
	if err != nil {
		return nil, err
	}

	capabilities, err := device.GetCapabilities(ctx, "All")
	if err != nil {
		return nil, err
	}

	media := loadMedia(capabilities, client)

	return &Services{
		Device: device,
		Media:  media,
	}, nil
}

// WithAuth ...
func WithAuth(username, password string) Option {
	return func(client *networking.Client) error {
		client.SetAuth(networking.ClientAuth{
			Username: username,
			Password: password,
		})
		return nil
	}
}

func loadMedia(capabilities *device.Capabilities, client *networking.Client) media.API {
	endpoint := string(capabilities.Media.XAddr)
	if len(endpoint) == 0 {
		return nil
	}
	client.AddEndpoint("media", endpoint)
	return media.New(client)
}
