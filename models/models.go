package models

import "encoding/xml"

type HediffDef struct {
	XMLName    xml.Name `xml:"HediffDef"`
	DefName    string   `xml:"defName"`
	ParentName string   `xml:"ParentName,attr"`
}

type ThingDef struct {
	XMLName    xml.Name `xml:"ThingDef"`
	DefName    string   `xml:"defName"`
	ParentName string   `xml:"ParentName,attr"`
}

type RecipeDef struct {
	XMLName    xml.Name `xml:"RecipeDef"`
	DefName    string   `xml:"defName"`
	ParentName string   `xml:"ParentName,attr"`
}

type Defs struct {
	Defs []any
}
