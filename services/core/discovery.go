// Package core ...
package core

import (
	wsdiscovery "github.com/jfsmig/onvif/ws-discovery"
	"github.com/korprulu/go-onvif-s/services"
)

// Discover send multicast discovery messages
func Discover(interfaceName string) ([]services.DeviceInfo, error) {
	clientInfos, err := wsdiscovery.GetAvailableDevicesAtSpecificEthernetInterface(interfaceName)
	if err != nil {
		return nil, err
	}

	infos := make([]services.DeviceInfo, len(clientInfos))
	for i, clientInfo := range clientInfos {
		infos[i] = services.DeviceInfo{
			Addr: clientInfo.Xaddr,
			UUID: clientInfo.Uuid,
		}
	}

	return infos, nil
}
