package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/korprulu/go-onvif-s/wsdd"
)

func main() {
	uuuid, _ := uuid.NewRandom()

	probe := wsdd.Probe{
		MessageID: "uuid:" + uuuid.String(),
		Types: []string{
			"dn:NetworkVideoTransmitter",
		},
		Namespace: map[string]string{
			"dn": "http://www.onvif.org/ver10/network/wsdl",
		},
	}
	soap := wsdd.NewProbeMessage(probe)

	if err := wsdd.SendProbe(soap); err != nil {
		fmt.Println(err)
	}
}
