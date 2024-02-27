package xdf2_models

import "encoding/xml"

type CodeListeReferenz struct {
	XmlName                  xml.Name                 `xml:"codelisteReferenz"`
	Identifikation           Identifikation           `xml:"identifikation"`
	GenericodeIdentification GenericodeIdentification `xml:"genericodeIdentification"`
}

type GenericodeIdentification struct {
	XmlName                 xml.Name `xml:"genericodeIdentification"`
	CanonicalIdentification string   `xml:"canonicalIdentification"`
	Version                 string   `xml:"version"`
	CanonicalVersionUri     string   `xml:"canonicalVersionUri"`
}
