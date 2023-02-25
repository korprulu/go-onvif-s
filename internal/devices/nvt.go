// Package devices ...
package devices

import (
	"github.com/jfsmig/onvif/networking"
)

// NVT ...
type NVT struct {
	client *networking.Client
}

// NewNVT returns a new NVT instance
func NewNVT(client *networking.Client) *NVT {
	return &NVT{
		client: client,
	}
}

// Info print the NVT info
func (nvt *NVT) Info() map[string]string {
	return map[string]string{
		"XAddr": nvt.client.GetEndpoint("device"),
		"UUID":  nvt.client.GetUUID(),
	}
}
