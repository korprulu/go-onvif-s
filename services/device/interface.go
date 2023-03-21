package device

import "context"

// API ...
type API interface {
	GetSystemDateAndTime(context.Context) (*SystemDateAndTime, error)
	GetCapabilities(context.Context, GetCapabilitiesInput) (*Capabilities, error)
}
