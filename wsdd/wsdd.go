package wsdd

import (
	"errors"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/korprulu/go-onvif-s/soap"
)

func SendProbe(probe *soap.Soap) ([]*ProbeMatches, error) {

	soapMsg, err := probe.Build()
	if err != nil {
		return nil, err
	}

	conn, err := net.ListenUDP("udp", nil)
	if err != nil {
		return nil, err
	}

	dstIP := &net.UDPAddr{
		IP:   net.IPv4(239, 255, 255, 250),
		Port: 3702,
	}

	if _, err := conn.WriteTo(soapMsg, dstIP); err != nil {
		return nil, err
	}

	if err := conn.SetReadDeadline(time.Now().Add(time.Second * 3)); err != nil {
		return nil, err
	}

	resp := make([][]byte, 0)

	buf := make([]byte, 8096)
	for {
		n, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			if errors.Is(err, os.ErrDeadlineExceeded) {
				break
			}
			fmt.Println(err)
		}
		resp = append(resp, buf[:n])
	}

	probeMatches := make([]*ProbeMatches, 0, len(resp))

	for _, r := range resp {
		if pm, err := ReadProbeMatches(r); err != nil {
			return nil, err
		} else {
			probeMatches = append(probeMatches, pm)
		}
	}

	return probeMatches, nil
}
