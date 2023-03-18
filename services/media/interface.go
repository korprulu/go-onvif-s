package media

import (
	"context"

	"github.com/jfsmig/onvif/xsd/onvif"
)

type MediaFunction interface {
	GetProfiles(context.Context) ([]Profile, error)
	GetVideoEncoderConfigurations(context.Context) ([]VideoEncoderConfiguration, error)
	GetGuaranteedNumberOfVideoEncoderInstances(context.Context, onvif.ReferenceToken) (*GuaranteedNumberOfVideoEncoderInstances, error)
}
