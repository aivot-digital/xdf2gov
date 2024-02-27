package xdf2_models

import "encoding/xml"

type Identifikation struct {
	XmlName xml.Name `xml:"identifikation"`
	Id      string   `xml:"id"`
	Version string   `xml:"version"`
}
