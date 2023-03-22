// Package media ...
package media

import (
	"context"

	"github.com/korprulu/go-onvif-s/internal/utils"
	mediaType "github.com/use-go/onvif/media"
	"github.com/use-go/onvif/sdk/media"
	onvifType "github.com/use-go/onvif/xsd/onvif"
)

// Profile ...
type Profile onvifType.Profile

// GetProfiles ...
func (m *Media) GetProfiles(ctx context.Context) ([]Profile, error) {
	resp, err := media.Call_GetProfiles(ctx, m.device, mediaType.GetProfiles{})
	if err != nil {
		return nil, err
	}

	result := utils.Map(resp.Profiles, func(p onvifType.Profile) Profile {
		return Profile(p)
	})

	return result, nil
}

type (
	// GetStreamURI ...
	GetStreamURI mediaType.GetStreamUri

	// StreamURI ...
	StreamURI mediaType.GetStreamUriResponse
)

// GetStreamURI ...
func (m *Media) GetStreamURI(ctx context.Context, input GetStreamURI) (*StreamURI, error) {
	resp, err := media.Call_GetStreamUri(ctx, m.device, mediaType.GetStreamUri(input))
	if err != nil {
		return nil, err
	}
	output := StreamURI(resp)
	return &output, nil
}
