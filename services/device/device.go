// Package device ...
package device

import (
	"context"

	"github.com/jfsmig/onvif/device"
	onvifDevice "github.com/jfsmig/onvif/device"
	"github.com/jfsmig/onvif/networking"
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

// GetSystemDateAndTimeResponse ...
type GetSystemDateAndTimeResponse device.GetSystemDateAndTimeResponse

// GetSystemDateAndTime ...
func (d *Device) GetSystemDateAndTime(ctx context.Context) (*GetSystemDateAndTimeResponse, error) {
	resp, err := onvifDevice.Call_GetSystemDateAndTime(ctx, d.client, onvifDevice.GetSystemDateAndTime{})
	if err != nil {
		return nil, err
	}

	result := GetSystemDateAndTimeResponse(resp)

	return &result, nil
}

type (
	// GetCapabilities ...
	GetCapabilities device.GetCapabilities

	// GetCapabilitiesResponse ...
	GetCapabilitiesResponse device.GetCapabilitiesResponse
)

// GetCapabilities ...
func (d *Device) GetCapabilities(ctx context.Context, input GetCapabilities) (*GetCapabilitiesResponse, error) {
	resp, err := onvifDevice.Call_GetCapabilities(ctx, d.client, device.GetCapabilities(input))
	if err != nil {
		return nil, err
	}

	result := GetCapabilitiesResponse(resp)

	return &result, nil
}
