// Package media ...
package media

import (
	"context"

	"github.com/jfsmig/onvif/media"
	"github.com/jfsmig/onvif/xsd/onvif"
)

// GetVideoEncoderConfigurationsResponse ...
type GetVideoEncoderConfigurationsResponse media.GetVideoEncoderConfigurationsResponse

// GetVideoEncoderConfigurations returns all of the video encoder configurations
func (m *Media) GetVideoEncoderConfigurations(ctx context.Context) (*GetVideoEncoderConfigurationsResponse, error) {
	resp, err := media.Call_GetVideoEncoderConfigurations(ctx, m.client, media.GetVideoEncoderConfigurations{})
	if err != nil {
		return nil, err
	}

	result := GetVideoEncoderConfigurationsResponse(resp)

	return &result, nil
}

// GetGuaranteedNumberOfVideoEncoderInstancesResponse ...
type GetGuaranteedNumberOfVideoEncoderInstancesResponse media.GetGuaranteedNumberOfVideoEncoderInstancesResponse

// GetGuaranteedNumberOfVideoEncoderInstances ...
func (m *Media) GetGuaranteedNumberOfVideoEncoderInstances(ctx context.Context, configurationToken onvif.ReferenceToken) (*GetGuaranteedNumberOfVideoEncoderInstancesResponse, error) {
	request := media.GetGuaranteedNumberOfVideoEncoderInstances{ConfigurationToken: configurationToken}
	resp, err := media.Call_GetGuaranteedNumberOfVideoEncoderInstances(ctx, m.client, request)
	if err != nil {
		return nil, err
	}

	result := GuaranteedNumberOfVideoEncoderInstances(resp)

	return &result, nil
}
