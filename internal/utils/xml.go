// Package utils ...
package utils

import (
	"github.com/go-xmlfmt/xmlfmt"
)

// XMLFormat formats XML content
func XMLFormat(xml []byte) string {
	return xmlfmt.FormatXML(string(xml), "", "  ", true)
}
