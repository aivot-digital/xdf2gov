package xdf2_models

import (
	"encoding/xml"
)

type Stammdatenschema struct {
	XmlName                                 xml.Name         `xml:"stammdatenschema"`
	Identifikation                          Identifikation   `xml:"identifikation"`
	Name                                    string           `xml:"name"`
	BezeichnungEingabe                      string           `xml:"bezeichnungEingabe"`
	BezeichnungAusgabe                      string           `xml:"bezeichnungAusgabe"`
	Beschreibung                            string           `xml:"beschreibung"`
	Definition                              string           `xml:"definition"`
	Bezug                                   string           `xml:"bezug"`
	Status                                  CodeListeWrapper `xml:"status"`
	FachlicherErsteller                     string           `xml:"fachlicherErsteller"`
	Veroeffentlichungsdatum                 string           `xml:"veroeffentlichungsdatum"`
	AbleitungsmodifikationenStruktur        CodeListeWrapper `xml:"ableitungsmodifikationenStruktur"`
	AbleitungsmodifikationenRepraesentation CodeListeWrapper `xml:"ableitungsmodifikationenRepraesentation"`
	Struktur                                []Struktur       `xml:"struktur"`
}
