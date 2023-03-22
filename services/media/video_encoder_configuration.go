// Package media ...
package media

import (
	"context"

	"github.com/korprulu/go-onvif-s/internal/utils"
	mediaType "github.com/use-go/onvif/media"
	"github.com/use-go/onvif/sdk/media"
	onvifType "github.com/use-go/onvif/xsd/onvif"
)

// VideoEncoderConfiguration ...
type VideoEncoderConfiguration onvifType.VideoEncoderConfiguration

// GetVideoEncoderConfigurations returns all of the video encoder configurations
func (m *Media) GetVideoEncoderConfigurations(ctx context.Context) ([]VideoEncoderConfiguration, error) {
	resp, err := media.Call_GetVideoEncoderConfigurations(ctx, m.device, mediaType.GetVideoEncoderConfigurations{})
	if err != nil {
		return nil, err
	}

	result := utils.Map(resp.Configurations, func(cfg onvifType.VideoEncoderConfiguration) VideoEncoderConfiguration {
		return VideoEncoderConfiguration(cfg)
	})

	return result, nil
}

type (
	// GetGuaranteedNumberOfVideoEncoderInstances ...
	GetGuaranteedNumberOfVideoEncoderInstances mediaType.GetGuaranteedNumberOfVideoEncoderInstances

	// GuaranteedNumberOfVideoEncoderInstances ...
	GuaranteedNumberOfVideoEncoderInstances mediaType.GetGuaranteedNumberOfVideoEncoderInstancesResponse
)

// GetGuaranteedNumberOfVideoEncoderInstances ...
func (m *Media) GetGuaranteedNumberOfVideoEncoderInstances(ctx context.Context, input GetGuaranteedNumberOfVideoEncoderInstances) (*GuaranteedNumberOfVideoEncoderInstances, error) {
	resp, err := media.Call_GetGuaranteedNumberOfVideoEncoderInstances(ctx, m.device, mediaType.GetGuaranteedNumberOfVideoEncoderInstances(input))
	if err != nil {
		return nil, err
	}

	result := GuaranteedNumberOfVideoEncoderInstances(resp)

	return &result, nil
}
