// Package devices ...
package devices

import (
	"context"
	"errors"
	"net/http"

	"github.com/jfsmig/onvif/networking"
	"github.com/jfsmig/onvif/sdk"
	"github.com/jfsmig/onvif/xsd/onvif"
)

// NVT ...
type NVT struct {
	device     sdk.Appliance
	client     *networking.Client
	descriptor *sdk.DeviceDescriptor
	profiles   *sdk.Profiles
	media      *sdk.Media
}

// NVTConfig ...
type NVTConfig struct {
	IPAddr     string
	UUID       string
	Username   string
	Password   string
	httpClient *http.Client
}

// NewNVT returns a new NVT instance
func NewNVT(ctx context.Context, cfg NVTConfig) (*NVT, error) {
	info := networking.ClientInfo{
		Xaddr: cfg.IPAddr,
		Uuid:  cfg.UUID,
	}

	auth := networking.ClientAuth{
		Username: cfg.Username,
		Password: cfg.Password,
	}

	client, err := networking.NewClient(info, nil)
	if err != nil {
		return nil, err
	}

	device, err := sdk.WrapClient(ctx, client, auth)
	if err != nil {
		return nil, err
	}

	return &NVT{
		device: device,
		client: client,
	}, nil
}

// Descriptor gets the device descriptor
func (nvt *NVT) Descriptor(ctx context.Context) *sdk.DeviceDescriptor {
	if nvt.descriptor != nil {
		return nvt.descriptor
	}
	deviceDescriptor := nvt.device.FetchDeviceDescriptor(ctx)
	nvt.descriptor = &deviceDescriptor
	return nvt.descriptor
}

// Profiles get device profiles
func (nvt *NVT) Profiles(ctx context.Context) *sdk.Profiles {
	if nvt.profiles != nil {
		return nvt.profiles
	}
	profiles := nvt.device.FetchProfiles(ctx)
	nvt.profiles = &profiles
	return nvt.profiles
}

// GetStreamURI ...
func (nvt *NVT) GetStreamURI(ctx context.Context, profileToken string) (string, error) {
	profiles := nvt.Profiles(ctx)

	if profile, ok := profiles.Profiles[onvif.ReferenceToken(profileToken)]; ok {
		return string(profile.Uris.Stream.Uri), nil
	}

	return "", errors.New("profile not found")
}

// GetMedia ...
func (nvt *NVT) GetMedia(ctx context.Context) (*sdk.Media, error) {
	if nvt.media != nil {
		return nvt.media, nil
	}

	media := nvt.device.FetchMedia(ctx)
	nvt.media = &media

	return &media, nil
}

// Initial ...
func (nvt *NVT) Initial(ctx context.Context) error {
	return nil
}
