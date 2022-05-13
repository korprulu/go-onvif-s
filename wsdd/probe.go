package wsdd

import (
	"fmt"
	"strings"

	"github.com/beevik/etree"
	"github.com/korprulu/go-onvif-s/soap"
)

type Probe struct {
	Namespace map[string]string
	MessageID string
	Scopes    []string
	Types     []string
}

func NewProbeMessage(probe Probe) *soap.Soap {

	actionHeader := etree.NewElement("w:Action").
		CreateAttr(soap.SoapEnvelopeNS+":mustUnderstand", "1").Element()
	actionHeader.SetText("http://schemas.xmlsoap.org/ws/2005/04/discovery/Probe")
	messageIDHeader := etree.NewElement("w:MessageID")
	messageIDHeader.SetText(probe.MessageID)
	toHeader := etree.NewElement("w:To").
		CreateAttr(soap.SoapEnvelopeNS+":mustUnderstand", "1").Element()
	toHeader.SetText("urn:schemas-xmlsoap-org:ws:2005:04:discovery")

	probeBody := etree.NewElement("d:Probe")

	if len(probe.Types) > 0 {
		probeBody.CreateElement("d:Types").SetText(strings.Join(probe.Types, " "))
	}

	if len(probe.Scopes) > 0 {
		probeBody.CreateElement("d:Scopes").SetText(strings.Join(probe.Scopes, " "))
	}

	return soap.New().
		AddGlobalNamespaces(map[string]string{
			"w": "http://schemas.xmlsoap.org/ws/2004/08/addressing",
			"d": "http://schemas.xmlsoap.org/ws/2005/04/discovery",
		}).
		AddGlobalNamespaces(probe.Namespace).
		AddHeaders(actionHeader, messageIDHeader, toHeader).
		AddBodies(probeBody)
}

type ProbeMatchesHeader struct {
	RelatesTo string
}

type ProbeMatches struct {
	Header ProbeMatchesHeader
}

func ReadProbeMatches(data []byte) (*ProbeMatches, error) {

	doc := etree.NewDocument()

	if err := doc.ReadFromBytes(data); err != nil {
		return nil, err
	}

	envelope := doc.FindElement("/Envelope[namespace-uri()='http://www.w3.org/2003/05/soap-envelope']")
	if envelope == nil {
		return nil, fmt.Errorf("read ProbeMatches failed: soap envelope not found")
	}

	header := envelope.FindElement("./Header[namespace-uri()='http://www.w3.org/2003/05/soap-envelope']")
	if header == nil {
		return nil, fmt.Errorf("read ProbeMatches failed: soap header not found")
	}

	body := envelope.FindElement("./Body[namespace-uri()='http://www.w3.org/2003/05/soap-envelope']")
	if body == nil {
		return nil, fmt.Errorf("read ProbeMatches failed: soap body not found")
	}

	relatesTo := header.FindElement("./RelatesTo[namespace-uri()='http://schemas.xmlsoap.org/ws/2004/08/addressing']")
	if relatesTo == nil {
		return nil, fmt.Errorf("read ProbeMatches failed: RelatesTo header not found")
	}

	return nil, nil
}
