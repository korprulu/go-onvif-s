// Package services ...
package services

import (
	"context"
	"net/http"

	"github.com/korprulu/go-onvif-s/services/device"
	"github.com/korprulu/go-onvif-s/services/media"
	"github.com/use-go/onvif"
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
	deviceParams := onvif.DeviceParams{
		Xaddr:      cfg.IPAddr,
		Username:   cfg.Username,
		Password:   cfg.Password,
		HttpClient: cfg.HTTPClient,
	}

	// the device endpoint will be added when create new client
	onvifDevice, err := onvif.NewDevice(deviceParams)
	if err != nil {
		return nil, err
	}

	dev, err := device.New(ctx, onvifDevice)
	if err != nil {
		return nil, err
	}

	capabilities, err := dev.GetCapabilities(ctx, device.GetCapabilitiesInput{Category: "All"})
	if err != nil {
		return nil, err
	}

	return &Services{
		Device:       dev,
		Media:        media.New(onvifDevice),
		rtpOverTCP:   bool(capabilities.Media.StreamingCapabilities.RTP_TCP),
		rtpMulticast: bool(capabilities.Media.StreamingCapabilities.RTPMulticast),
		rtsp:         bool(capabilities.Media.StreamingCapabilities.RTP_RTSP_TCP),
	}, nil
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
