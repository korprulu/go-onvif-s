// Package media ...
package media

import (
	"context"

	"github.com/jfsmig/onvif/media"
	"github.com/jfsmig/onvif/xsd/onvif"
	"github.com/korprulu/go-onvif-s/internal/utils"
)

// Profile ...
type Profile onvif.Profile

// GetProfiles ...
func (m *Media) GetProfiles(ctx context.Context) ([]Profile, error) {
	resp, err := media.Call_GetProfiles(ctx, m.client, media.GetProfiles{})
	if err != nil {
		return nil, err
	}

	profiles := utils.Map(resp.Profiles, func(p onvif.Profile) Profile { return Profile(p) })

	return profiles, nil
}
