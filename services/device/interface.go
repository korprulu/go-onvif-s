package device

import "context"

type DeviceFunction interface {
	GetSystemDateAndTime(ctx context.Context) (*SystemDateAndTime, error)
	GetCapabilities(ctx context.Context, category string) (*Capabilities, error)
}
