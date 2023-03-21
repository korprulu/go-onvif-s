// Package media ...
package media

import (
	"context"
	"fmt"
	"io"

	"github.com/jfsmig/onvif/media"
	"github.com/jfsmig/onvif/xsd/onvif"
	"github.com/korprulu/go-onvif-s/internal/utils"
)

// Profile ...
type Profile onvif.Profile

// GetProfiles ...
func (m *Media) GetProfiles(ctx context.Context) ([]Profile, error) {
	rawResp, err := m.client.CallMethod(ctx, media.GetProfiles{})
	if err != nil {
		return nil, err
	}
	defer rawResp.Body.Close()

	payload, err := io.ReadAll(rawResp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(utils.XMLFormat(payload))

	resp, err := media.Call_GetProfiles(ctx, m.client, media.GetProfiles{})
	if err != nil {
		return nil, err
	}

	result := utils.Map(resp.Profiles, func(p onvif.Profile) Profile {
		return Profile(p)
	})

	return result, nil
}

type (
	// GetStreamURI ...
	GetStreamURI media.GetStreamUri

	// StreamURI ...
	StreamURI media.GetStreamUriResponse
)

// GetStreamURI ...
func (m *Media) GetStreamURI(ctx context.Context, input GetStreamURI) (*StreamURI, error) {
	resp, err := media.Call_GetStreamUri(ctx, m.client, media.GetStreamUri(input))
	if err != nil {
		return nil, err
	}
	output := StreamURI(resp)
	return &output, nil
}
