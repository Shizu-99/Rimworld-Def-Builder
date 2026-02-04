package models

import "encoding/xml"

type BodyPart struct {
	Hediff  HediffDef
	Thing   ThingDef
	Install RecipeDef
	Remove  RecipeDef
}

type HediffDef struct {
	DefName    string `xml:"defName"`
	ParentName string `xml:"ParentName,attr"`
}

type ThingDef struct {
	DefName    string `xml:"defName"`
	ParentName string `xml:"ParentName,attr"`
}

type RecipeDef struct {
	DefName    string `xml:"defName"`
	ParentName string `xml:"ParentName,attr"`
}

type Defs struct {
	XMLName xml.Name `xml:"Defs"`
	defs    []BodyPart
}
