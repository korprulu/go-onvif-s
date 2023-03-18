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
	req := media.GetVideoEncoderConfigurations{}
	resp, err := media.Call_GetVideoEncoderConfigurations(ctx, m.client, req)
	if err != nil {
		return nil, err
	}

	configs := utils.Map(resp.Configurations, func(cfg onvif.VideoEncoderConfiguration) VideoEncoderConfiguration {
		return VideoEncoderConfiguration(cfg)
	})

	return configs, nil
}

// GuaranteedNumberOfVideoEncoderInstances ...
type GuaranteedNumberOfVideoEncoderInstances media.GetGuaranteedNumberOfVideoEncoderInstancesResponse

// GetGuaranteedNumberOfVideoEncoderInstances ...
func (m *Media) GetGuaranteedNumberOfVideoEncoderInstances(ctx context.Context, configurationToken onvif.ReferenceToken) (*GuaranteedNumberOfVideoEncoderInstances, error) {
	request := media.GetGuaranteedNumberOfVideoEncoderInstances{ConfigurationToken: configurationToken}
	resp, err := media.Call_GetGuaranteedNumberOfVideoEncoderInstances(ctx, m.client, request)
	if err != nil {
		return nil, err
	}

	result := GuaranteedNumberOfVideoEncoderInstances(resp)

	return &result, nil
}
