package xrepository_models

import "encoding/xml"

type ColumnSet struct {
	XmlName xml.Name     `xml:"ColumnSet"`
	Key     ColumnSetKey `xml:"Key"`
}

type ColumnSetKey struct {
	XmlName   xml.Name        `xml:"Key"`
	ColumnRef ColumnSetKeyRef `xml:"ColumnRef"`
}

type ColumnSetKeyRef struct {
	XmlName xml.Name `xml:"ColumnRef"`
	Ref     string   `xml:"Ref,attr"`
}
