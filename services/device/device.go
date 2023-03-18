// Package device ...
package device

import (
	"context"

	onvifDevice "github.com/jfsmig/onvif/device"
	"github.com/jfsmig/onvif/networking"
	"github.com/jfsmig/onvif/xsd/onvif"
)

// Device ...
type Device struct {
	client *networking.Client
}

var _ DeviceFunction = (*Device)(nil)

func New(ctx context.Context, client *networking.Client) (DeviceFunction, error) {
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

// Capabilities ...
type Capabilities onvif.Capabilities

// GetCapabilities ...
func (d *Device) GetCapabilities(ctx context.Context, category string) (*Capabilities, error) {
	req := onvifDevice.GetCapabilities{Category: onvif.CapabilityCategory(category)}
	resp, err := onvifDevice.Call_GetCapabilities(ctx, d.client, req)
	if err != nil {
		return nil, err
	}

	result := Capabilities(resp.Capabilities)

	return &result, nil
}
