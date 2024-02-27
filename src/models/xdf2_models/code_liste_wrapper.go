package xdf2_models

type CodeListeWrapper struct {
	ListURI       string `xml:"listURI,attr"`
	ListVersionID string `xml:"listVersionID,attr"`
	Code          string `xml:"code"`
}
