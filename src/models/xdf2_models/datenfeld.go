package xdf2_models

import "encoding/xml"

type Datenfeld struct {
	XmlName             xml.Name          `xml:"datenfeld"`
	Identifikation      Identifikation    `xml:"identifikation"`
	Name                string            `xml:"name"`
	BezeichnungEingabe  string            `xml:"bezeichnungEingabe"`
	BezeichnungAusgabe  string            `xml:"bezeichnungAusgabe"`
	Beschreibung        string            `xml:"beschreibung"`
	Definition          string            `xml:"definition"`
	Bezug               string            `xml:"bezug"`
	Status              CodeListeWrapper  `xml:"status"`
	FachlicherErsteller string            `xml:"fachlicherErsteller"`
	Schemaelementart    CodeListeWrapper  `xml:"schemaelementart"`
	HilfetextEingabe    string            `xml:"hilfetextEingabe"`
	HilfetextAusgabe    string            `xml:"hilfetextAusgabe"`
	Feldart             CodeListeWrapper  `xml:"feldart"`
	Datentyp            CodeListeWrapper  `xml:"datentyp"`
	Praezisierung       string            `xml:"praezisierung"`
	Inhalt              string            `xml:"inhalt"`
	CodelisteReferenz   CodeListeReferenz `xml:"codelisteReferenz"`
}
