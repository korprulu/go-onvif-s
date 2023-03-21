// Package media ...
package media

import "github.com/jfsmig/onvif/networking"

// Media ...
type Media struct {
	client *networking.Client
}

var _ API = (*Media)(nil)

// New ...
func New(client *networking.Client) *Media {
	return &Media{client: client}
}
