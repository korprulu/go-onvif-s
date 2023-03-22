// Package device ...
package device

import (
	"context"

	"github.com/use-go/onvif"
	deviceType "github.com/use-go/onvif/device"
	"github.com/use-go/onvif/sdk/device"
	onvifType "github.com/use-go/onvif/xsd/onvif"
)

// Device ...
type Device struct {
	device *onvif.Device
}

var _ API = (*Device)(nil)

// New ...
func New(ctx context.Context, device *onvif.Device) (*Device, error) {
	dev := &Device{device: device}

	if _, err := dev.GetSystemDateAndTime(ctx); err != nil {
		return nil, err
	}

	return dev, nil
}

// SystemDateAndTime ...
type SystemDateAndTime onvifType.SystemDateTime

// GetSystemDateAndTime ...
func (d *Device) GetSystemDateAndTime(ctx context.Context) (*SystemDateAndTime, error) {
	resp, err := device.Call_GetSystemDateAndTime(ctx, d.device, deviceType.GetSystemDateAndTime{})
	if err != nil {
		return nil, err
	}

	result := SystemDateAndTime(resp.SystemDateAndTime)

	return &result, nil
}

type (
	// GetCapabilitiesInput ...
	GetCapabilitiesInput deviceType.GetCapabilities

	// Capabilities ...
	Capabilities onvifType.Capabilities
)

// GetCapabilities ...
func (d *Device) GetCapabilities(ctx context.Context, input GetCapabilitiesInput) (*Capabilities, error) {
	resp, err := device.Call_GetCapabilities(ctx, d.device, deviceType.GetCapabilities(input))
	if err != nil {
		return nil, err
	}

	result := Capabilities(resp.Capabilities)

	return &result, nil
}
