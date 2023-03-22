// Package media ...
package media

import (
	"github.com/use-go/onvif"
)

// Media ...
type Media struct {
	device *onvif.Device
}

var _ API = (*Media)(nil)

// New ...
func New(device *onvif.Device) *Media {
	return &Media{device: device}
}
