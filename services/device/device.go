// Package device ...
package device

import (
	"context"
	"net/http"

	onvifDevice "github.com/jfsmig/onvif/device"
	"github.com/jfsmig/onvif/networking"
	"github.com/jfsmig/onvif/xsd/onvif"
)

// Device ...
type Device struct {
	client       *networking.Client
	capabilities *Capabilities
}

// Option ...
type Option func(*Device) error

// Info ...
type Info struct {
	Addr string
	UUID string
}

// New creates a new device
func New(ctx context.Context, info Info, httpClient *http.Client, options ...Option) (*Device, error) {
	clientInfo := networking.ClientInfo{
		Xaddr: info.Addr,
		Uuid:  info.UUID,
	}

	client, err := networking.NewClient(clientInfo, httpClient)
	if err != nil {
		return nil, err
	}

	device := &Device{client: client}

	for _, opt := range options {
		if err := opt(device); err != nil {
			return nil, err
		}
	}

	if _, err := device.GetSystemDateAndTime(ctx); err != nil {
		return nil, err
	}

	capabilities, err := device.GetCapabilities(ctx, "All")
	if err != nil {
		return nil, err
	}

	device.capabilities = capabilities

	return device, nil
}

// WithAuth ...
func WithAuth(username, password string) Option {
	return func(dev *Device) error {
		dev.client.SetAuth(networking.ClientAuth{
			Username: username,
			Password: password,
		})
		return nil
	}
}

// SystemDateAndTime ...
type SystemDateAndTime onvif.SystemDateTime

// GetSystemDateAndTime ...
func (d *Device) GetSystemDateAndTime(ctx context.Context) (*SystemDateAndTime, error) {
	resp, err := onvifDevice.Call_GetSystemDateAndTime(ctx, d.client, onvifDevice.GetSystemDateAndTime{})
	if err != nil {
		return nil, err
	}

	result := SystemDateAndTime(resp.SystemDateAndTime)

	return &result, nil
}

// Capabilities ...
type Capabilities onvif.Capabilities

// GetCapabilities ...
func (d *Device) GetCapabilities(ctx context.Context, category string) (*Capabilities, error) {
	if d.capabilities != nil {
		return d.capabilities, nil
	}

	req := onvifDevice.GetCapabilities{Category: onvif.CapabilityCategory(category)}
	resp, err := onvifDevice.Call_GetCapabilities(ctx, d.client, req)
	if err != nil {
		return nil, err
	}

	result := Capabilities(resp.Capabilities)

	return &result, nil
}
