package device

import "context"

// API ...
type API interface {
	GetSystemDateAndTime(context.Context) (*GetSystemDateAndTimeResponse, error)
	GetCapabilities(context.Context, GetCapabilities) (*GetCapabilitiesResponse, error)
}
