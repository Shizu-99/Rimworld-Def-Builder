package models

import "encoding/xml"

type HediffDef struct {
	XMLName               xml.Name  `xml:"HediffDef"`
	DefName               string    `xml:"defName"`
	ParentName            string    `xml:"ParentName,attr"`
	Label                 string    `xml:"label"`
	LabelNoun             string    `xml:"labelNoun"`
	Description           string    `xml:"description"`
	DescriptionHyperlinks string    `xml:"descriptionHyperlinks>ThingDef"`
	SpawnThingOnRemoved   string    `xml:"spawnThingOnRemoved"`
	AddedPartProps        PartProps `xml:"addedPartProps,omitempty"`
	Stages                Stages    `xml:"stages>li,omitempty"`
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

type PartProps struct {
	Solid bool `xml:"solid,omitempty"`
}

type StatOffsets struct {
	XMLName                 xml.Name `xml:"statOffsets"`
	ConstructionSpeedFactor float32  `xml:"ConstructionSpeedFactor,omitempty"`
	MiningSpeed             float32  `xml:"MiningSpeed,omitempty"`
}

type CapMods struct {
	XMLName    xml.Name   `xml:"capMods"`
	Capacities []Capacity `xml:"li"`
}

type Capacity struct {
	Capacity string  `xml:"capacity"`
	Offset   float32 `xml:"offset"`
}

type Stages struct {
	Stats        StatOffsets
	CapacityMods CapMods
}
