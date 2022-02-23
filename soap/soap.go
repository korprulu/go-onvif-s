package soap

import (
	"github.com/beevik/etree"
)

type Soap struct {
	document *etree.Document
	header   *etree.Element
	body     *etree.Element
}

const SoapEnvelopeNS = "e"

var DefaultGlobalNamespaces = map[string]string{
	SoapEnvelopeNS: "http://www.w3.org/2003/05/soap-envelope",
}

func New() *Soap {
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)

	envelope := doc.CreateElement(SoapEnvelopeNS + ":Envelope")

	header := envelope.CreateElement(SoapEnvelopeNS + ":Header")
	body := envelope.CreateElement(SoapEnvelopeNS + ":Body")

	s := &Soap{
		document: doc,
		header:   header,
		body:     body,
	}

	return s.AddGlobalNamespaces(DefaultGlobalNamespaces)
}

func (s *Soap) AddGlobalNamespaces(ns map[string]string) *Soap {
	for k, v := range ns {
		s.document.Root().CreateAttr("xmlns:"+k, v)
	}
	return s
}

func (s *Soap) AddHeaders(headers ...etree.Token) *Soap {
	for _, h := range headers {
		s.header.AddChild(h)
	}
	return s
}

func (s *Soap) AddBodies(bodies ...etree.Token) *Soap {
	for _, b := range bodies {
		s.body.AddChild(b)
	}
	return s
}

func (s *Soap) Build() ([]byte, error) {
	return s.document.WriteToBytes()
}
