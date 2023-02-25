// Package discovery ...
package discovery

import (
	"github.com/jfsmig/onvif/networking"
	wsdiscovery "github.com/jfsmig/onvif/ws-discovery"
	"github.com/korprulu/go-onvif-s/internal/devices"
)

// Probe send multicast discovery messages
func Probe(interfaceName string) ([]*devices.NVT, error) {

	clientInfos, err := wsdiscovery.GetAvailableDevicesAtSpecificEthernetInterface(interfaceName)
	if err != nil {
		return nil, err
	}

	nvts := make([]*devices.NVT, len(clientInfos))
	for i, clientInfo := range clientInfos {
		client, err := networking.NewClient(clientInfo, nil)
		if err != nil {
			return nil, err
		}
		nvts[i] = devices.NewNVT(client)
	}

	return nvts, nil
}
