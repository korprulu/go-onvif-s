package media

import (
	"context"
)

// API ...
type API interface {
	GetProfiles(context.Context) ([]Profile, error)
	GetVideoEncoderConfigurations(context.Context) ([]VideoEncoderConfiguration, error)
	GetGuaranteedNumberOfVideoEncoderInstances(context.Context, GetGuaranteedNumberOfVideoEncoderInstances) (*GuaranteedNumberOfVideoEncoderInstances, error)
	GetStreamURI(context.Context, GetStreamURI) (*StreamURI, error)
}
