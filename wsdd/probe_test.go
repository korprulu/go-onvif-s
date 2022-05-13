package wsdd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProbeMessageBuilder(t *testing.T) {
	probe := Probe{
		MessageID: "uuid:abc123",
		Scopes:    []string{"test:aaa", "test:bbb"},
		Types:     []string{"test:111", "test:222"},
		Namespace: map[string]string{
			"test": "http://testing.com",
		},
	}
	soap := NewProbeMessage(probe)

	content, err := soap.Build()
	assert.Nil(t, err)

	expected := `<?xml version="1.0" encoding="UTF-8"?><e:Envelope xmlns:e="http://www.w3.org/2003/05/soap-envelope" xmlns:w="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:d="http://schemas.xmlsoap.org/ws/2005/04/discovery" xmlns:test="http://testing.com"><e:Header><w:Action e:mustUnderstand="1">http://schemas.xmlsoap.org/ws/2005/04/discovery/Probe</w:Action><w:MessageID>uuid:abc123</w:MessageID><w:To e:mustUnderstand="1">urn:schemas-xmlsoap-org:ws:2005:04:discovery</w:To></e:Header><e:Body><d:Probe><d:Types>test:111 test:222</d:Types><d:Scopes>test:aaa test:bbb</d:Scopes></d:Probe></e:Body></e:Envelope>`

	assert.Equal(t, expected, string(content))
}
