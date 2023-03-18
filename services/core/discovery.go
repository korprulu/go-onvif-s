// Package core ...
package core

import (
	wsdiscovery "github.com/jfsmig/onvif/ws-discovery"
	"github.com/korprulu/go-onvif-s/services/device"
)

// Discover send multicast discovery messages
func Discover(interfaceName string) ([]device.Info, error) {
	clientInfos, err := wsdiscovery.GetAvailableDevicesAtSpecificEthernetInterface(interfaceName)
	if err != nil {
		return nil, err
	}

	infos := make([]device.Info, len(clientInfos))
	for i, clientInfo := range clientInfos {
		infos[i] = device.Info{
			Addr: clientInfo.Xaddr,
			UUID: clientInfo.Uuid,
		}
	}

	return infos, nil
}
