package xrepository_models

import "encoding/xml"

type CodeList struct {
	XmlName        xml.Name       `xml:"CodeList"`
	ColumnSet      ColumnSet      `xml:"ColumnSet"`
	SimpleCodeList SimpleCodeList `xml:"SimpleCodeList"`
}
