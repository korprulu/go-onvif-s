// Package media ...
package media

import (
	"context"

	"github.com/jfsmig/onvif/media"
	"github.com/jfsmig/onvif/xsd/onvif"
	"github.com/korprulu/go-onvif-s/internal/utils"
)

// VideoEncoderConfiguration ...
type VideoEncoderConfiguration onvif.VideoEncoderConfiguration

// GetVideoEncoderConfigurations returns all of the video encoder configurations
func (m *Media) GetVideoEncoderConfigurations(ctx context.Context) ([]VideoEncoderConfiguration, error) {
	resp, err := media.Call_GetVideoEncoderConfigurations(ctx, m.client, media.GetVideoEncoderConfigurations{})
	if err != nil {
		return nil, err
	}

	result := utils.Map(resp.Configurations, func(cfg onvif.VideoEncoderConfiguration) VideoEncoderConfiguration {
		return VideoEncoderConfiguration(cfg)
	})

	return result, nil
}

type (
	// GetGuaranteedNumberOfVideoEncoderInstances ...
	GetGuaranteedNumberOfVideoEncoderInstances media.GetGuaranteedNumberOfVideoEncoderInstances

	// GuaranteedNumberOfVideoEncoderInstances ...
	GuaranteedNumberOfVideoEncoderInstances media.GetGuaranteedNumberOfVideoEncoderInstancesResponse
)

// GetGuaranteedNumberOfVideoEncoderInstances ...
func (m *Media) GetGuaranteedNumberOfVideoEncoderInstances(ctx context.Context, input GetGuaranteedNumberOfVideoEncoderInstances) (*GuaranteedNumberOfVideoEncoderInstances, error) {
	resp, err := media.Call_GetGuaranteedNumberOfVideoEncoderInstances(ctx, m.client, media.GetGuaranteedNumberOfVideoEncoderInstances(input))
	if err != nil {
		return nil, err
	}

	result := GuaranteedNumberOfVideoEncoderInstances(resp)

	return &result, nil
}
