package xdf2_models

import "encoding/xml"

type Struktur struct {
	XmlName  xml.Name `xml:"struktur"`
	Anzahl   string   `xml:"anzahl"`
	Bezug    string   `xml:"bezug"`
	Enthaelt Enthaelt `xml:"enthaelt"`
}
