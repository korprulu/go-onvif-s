// Package device ...
package device

import (
	"context"

	"github.com/jfsmig/onvif/device"
	onvifDevice "github.com/jfsmig/onvif/device"
	"github.com/jfsmig/onvif/networking"
	"github.com/jfsmig/onvif/xsd/onvif"
)

// Device ...
type Device struct {
	client *networking.Client
}

var _ API = (*Device)(nil)

// New ...
func New(ctx context.Context, client *networking.Client) (*Device, error) {
	dev := &Device{client: client}

	if _, err := dev.GetSystemDateAndTime(ctx); err != nil {
		return nil, err
	}

	return dev, nil
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

type (
	// GetCapabilitiesInput ...
	GetCapabilitiesInput device.GetCapabilities

	// Capabilities ...
	Capabilities onvif.Capabilities
)

// GetCapabilities ...
func (d *Device) GetCapabilities(ctx context.Context, input GetCapabilitiesInput) (*Capabilities, error) {
	resp, err := onvifDevice.Call_GetCapabilities(ctx, d.client, device.GetCapabilities(input))
	if err != nil {
		return nil, err
	}

	result := Capabilities(resp.Capabilities)

	return &result, nil
}
