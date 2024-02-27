package xdf2_models

import (
	"encoding/xml"
)

type Header struct {
	XmlName              xml.Name `xml:"header"`
	NachrichtenId        string   `xml:"nachrichtID"`
	Erstellungszeitpunkt string   `xml:"erstellungszeitpunkt"`
}
