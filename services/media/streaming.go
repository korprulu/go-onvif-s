// Package media ...
package media

import (
	"context"

	"github.com/jfsmig/onvif/media"
)

// GetProfilesResponse ...
type GetProfilesResponse media.GetProfilesResponse

// GetProfiles ...
func (m *Media) GetProfiles(ctx context.Context) (*GetProfilesResponse, error) {
	resp, err := media.Call_GetProfiles(ctx, m.client, media.GetProfiles{})
	if err != nil {
		return nil, err
	}

	result := GetProfilesResponse(resp)

	return &result, nil
}

type (
	// GetStreamURI ...
	GetStreamURI media.GetStreamUri

	// GetStreamURIResponse ...
	GetStreamURIResponse media.GetStreamUriResponse
)

// GetStreamURI ...
func (m *Media) GetStreamURI(ctx context.Context, input GetStreamURI) (*GetStreamURIResponse, error) {
	resp, err := media.Call_GetStreamUri(ctx, m.client, media.GetStreamUri(input))
	if err != nil {
		return nil, err
	}
	output := GetStreamURIResponse(resp)
	return &output, nil
}
