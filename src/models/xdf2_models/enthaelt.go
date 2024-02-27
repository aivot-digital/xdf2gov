package xdf2_models

import "encoding/xml"

type Enthaelt struct {
	XmlName         xml.Name        `xml:"enthaelt"`
	Datenfeldgruppe Datenfeldgruppe `xml:"datenfeldgruppe"`
	Datenfeld       Datenfeld       `xml:"datenfeld"`
}
