package xdf2_models

import "encoding/xml"

type Stammdatenschema0102 struct {
	XmlName          xml.Name         `xml:"xdatenfelder.stammdatenschema.0102"`
	Header           Header           `xml:"header"`
	Stammdatenschema Stammdatenschema `xml:"stammdatenschema"`
}
