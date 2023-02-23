package discovery

import (
	"errors"
	"net"
	"os"
	"strings"
	"time"

	"github.com/beevik/etree"
	"github.com/google/uuid"
	"github.com/korprulu/go-onvif-s/internal/devices"
	"github.com/korprulu/go-onvif-s/internal/soap"
	"github.com/rs/zerolog/log"
)

type Prober struct {
	messageID string
}

func (p *Prober) Probe() ([]*devices.NVT, error) {

	theUUID, _ := uuid.NewRandom()
	p.messageID = "uuid:" + theUUID.String()

	soapMsg, err := p.buildSoapMsg().Build()
	if err != nil {
		return nil, err
	}
	log.Debug().Msgf("probe message: %s", string(soapMsg))

	conn, err := net.ListenUDP("udp4", nil)
	if err != nil {
		return nil, err
	}
	log.Debug().Msgf("local addr: %v", conn.LocalAddr())

	if err := conn.SetReadDeadline(time.Now().Add(time.Second * 5)); err != nil {
		return nil, err
	}

	dstIP, err := net.ResolveUDPAddr("udp", "239.255.255.250:3702")
	if err != nil {
		return nil, err
	}

	if _, err := conn.WriteTo(soapMsg, dstIP); err != nil {
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
			return nil, err
		}

		if n == 0 {
			continue
		}

		dst := make([]byte, n)
		copy(dst, buf[:n])

		log.Debug().Msgf("received message: %s", dst)

		resp = append(resp, dst)
	}

	devices := make([]*devices.NVT, 0, len(resp))

	for _, r := range resp {
		if device, err := extractDevice(p.messageID, r); err != nil {
			return nil, err
		} else {
			devices = append(devices, device)
		}
	}

	return devices, nil
}

func (p *Prober) buildSoapMsg() *soap.Soap {

	namespaces := map[string]string{
		"w":  "http://schemas.xmlsoap.org/ws/2004/08/addressing",
		"d":  "http://schemas.xmlsoap.org/ws/2005/04/discovery",
		"dn": "http://www.onvif.org/ver10/network/wsdl",
	}

	actionHeader := etree.NewElement("w:Action").
		CreateAttr(soap.SoapEnvelopeNS+":mustUnderstand", "1").Element().
		CreateText("http://schemas.xmlsoap.org/ws/2005/04/discovery/Probe").Parent()

	messageIDHeader := etree.NewElement("w:MessageID").CreateText(p.messageID).Parent()

	toHeader := etree.NewElement("w:To").
		CreateAttr(soap.SoapEnvelopeNS+":mustUnderstand", "1").Element().
		CreateText("urn:schemas-xmlsoap-org:ws:2005:04:discovery").Parent()

	probeBody := etree.NewElement("d:Probe")
	probeBody.CreateElement("d:Types").CreateText("dn:NetworkVideoTransmitter")

	return soap.New().
		AddGlobalNamespaces(namespaces).
		AddHeaders(actionHeader, messageIDHeader, toHeader).
		AddBodies(probeBody)
}

func extractDevice(messageID string, probeMatches []byte) (*devices.NVT, error) {

	doc := etree.NewDocument()

	if err := doc.ReadFromBytes(probeMatches); err != nil {
		return nil, err
	}

	envelope := doc.FindElement("/Envelope[namespace-uri()='http://www.w3.org/2003/05/soap-envelope']")
	if envelope == nil {
		return nil, errors.New("read ProbeMatches failed: soap envelope not found")
	}

	header := envelope.FindElement("./Header[namespace-uri()='http://www.w3.org/2003/05/soap-envelope']")
	if header == nil {
		return nil, errors.New("read ProbeMatches failed: soap header not found")
	}

	relatesTo := header.FindElement("./RelatesTo[namespace-uri()='http://schemas.xmlsoap.org/ws/2004/08/addressing']")
	if relatesTo == nil {
		return nil, errors.New("read ProbeMatches failed: RelatesTo header not found")
	} else if relatesTo.Text() != messageID {
		return nil, errors.New("mismatched message id")
	}

	body := envelope.FindElement("./Body[namespace-uri()='http://www.w3.org/2003/05/soap-envelope']")
	if body == nil {
		return nil, errors.New("read ProbeMatches failed: soap body not found")
	}

	nvt := devices.NewNVT()

	xaddrs := body.FindElement("./ProbeMatches[namespace-uri()='http://schemas.xmlsoap.org/ws/2005/04/discovery']" +
		"/ProbeMatch[namespace-uri()='http://schemas.xmlsoap.org/ws/2005/04/discovery']" +
		"/XAddrs[namespace-uri()='http://schemas.xmlsoap.org/ws/2005/04/discovery']")

	nvt.Address = xaddrs.Text()

	scopes := body.FindElement("./ProbeMatches[namespace-uri()='http://schemas.xmlsoap.org/ws/2005/04/discovery']" +
		"/ProbeMatch[namespace-uri()='http://schemas.xmlsoap.org/ws/2005/04/discovery']" +
		"/Scopes[namespace-uri()='http://schemas.xmlsoap.org/ws/2005/04/discovery']")

	for _, scope := range strings.Split(scopes.Text(), " ") {
		scope = strings.TrimPrefix(scope, "onvif://www.onvif.org/")
		if len(scope) == 0 {
			continue
		}
		theFirstSlashIndex := strings.Index(scope, "/")
		nvt.Scopes[scope[0:theFirstSlashIndex]] = scope[theFirstSlashIndex+1:]
	}

	return nvt, nil
}
