// Package devices ...
package devices

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/jfsmig/onvif/device"
	"github.com/jfsmig/onvif/networking"
	"github.com/jfsmig/onvif/sdk"
	"github.com/korprulu/go-onvif-s/internal/utils"
)

// NVT ...
type NVT struct {
	device sdk.Appliance
}

// NVTConfig ...
type NVTConfig struct {
	IPAddr     string
	UUID       string
	Username   string
	Password   string
	httpClient *http.Client
}

// NewNVT returns a new NVT instance
func NewNVT(ctx context.Context, cfg NVTConfig) (*NVT, error) {
	info := networking.ClientInfo{
		Xaddr: cfg.IPAddr,
		Uuid:  cfg.UUID,
	}

	auth := networking.ClientAuth{
		Username: "admin",
		Password: "590885",
	}

	device, err := sdk.NewDevice(ctx, info, auth, cfg.httpClient)
	if err != nil {
		return nil, err
	}

	return &NVT{
		device: device,
	}, nil
}

// Info print the NVT info
func (nvt *NVT) Info() map[string]string {
	return map[string]string{
		"XAddr": nvt.client.GetEndpoint("device"),
		"UUID":  nvt.client.GetUUID(),
	}
}

// Initial ...
func (nvt *NVT) Initial(ctx context.Context) error {
	resp, err := nvt.client.CallMethod(ctx, &device.GetServices{})
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	respData, err := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		fmt.Println(utils.XMLFormat(respData))
		return fmt.Errorf("http request error: %d", resp.StatusCode)
	}

	fmt.Println(utils.XMLFormat(respData))

	return nil
}
