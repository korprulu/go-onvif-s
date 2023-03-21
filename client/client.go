// Package client ...
package client

import (
	"context"

	"github.com/korprulu/go-onvif-s/internal/utils"
	"github.com/korprulu/go-onvif-s/services"
)

// Client ...
type Client struct {
	svc *services.Services
}

// New creates a new client instance
func New(ctx context.Context, svc *services.Services) (*Client, error) {
	c := &Client{svc: svc}
	profiles, err := c.svc.Media.GetProfiles(ctx)
	if err != nil {
		return nil, err
	}
	for _, profile := range profiles {
		utils.PrintJSON(profile)
	}
	return c, nil
}
