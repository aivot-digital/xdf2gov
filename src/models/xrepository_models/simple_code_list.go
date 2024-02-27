package xrepository_models

import "encoding/xml"

type SimpleCodeList struct {
	XmlName xml.Name            `xml:"SimpleCodeList"`
	Rows    []SimpleCodeListRow `xml:"Row"`
}

type SimpleCodeListRow struct {
	XmlName xml.Name                 `xml:"Row"`
	Values  []SimpleCodeListRowValue `xml:"Value"`
}

type SimpleCodeListRowValue struct {
	XmlName     xml.Name `xml:"Value"`
	ColumnRef   string   `xml:"ColumnRef,attr"`
	SimpleValue string   `xml:"SimpleValue"`
}
