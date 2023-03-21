package media

import (
	"context"
)

// API ...
type API interface {
	GetProfiles(context.Context) (*GetProfilesResponse, error)
	GetVideoEncoderConfigurations(context.Context) (*GetVideoEncoderConfigurationsResponse, error)
	GetGuaranteedNumberOfVideoEncoderInstances(context.Context, GetGuaranteedNumberOfVideoEncoderInstances) (*GetGuaranteedNumberOfVideoEncoderInstancesResponse, error)
	GetStreamURI(context.Context, GetStreamURI) (*GetStreamURIResponse, error)
}
