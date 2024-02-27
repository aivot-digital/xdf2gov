package xdf2_models

import "encoding/xml"

type Regel struct {
	XmlName             xml.Name         `xml:"regel"`
	Identifikation      Identifikation   `xml:"identifikation"`
	Name                string           `xml:"name"`
	BezeichnungEingabe  string           `xml:"bezeichnungEingabe"`
	Beschreibung        string           `xml:"beschreibung"`
	Definition          string           `xml:"definition"`
	Bezug               string           `xml:"bezug"`
	Status              CodeListeWrapper `xml:"status"`
	FachlicherErsteller string           `xml:"fachlicherErsteller"`
	Script              string           `xml:"script"`
}
