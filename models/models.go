package models

import "encoding/xml"

type HediffDef struct {
	XMLName               xml.Name              `xml:"HediffDef"`
	DefName               string                `xml:"defName"`
	ParentName            string                `xml:"ParentName,attr"`
	Label                 string                `xml:"label"`
	LabelNoun             string                `xml:"labelNoun,omitempty"`
	Description           string                `xml:"description"`
	DescriptionHyperlinks DescriptionHyperlinks `xml:"descriptionHyperlinks"`
	SpawnThingOnRemoved   string                `xml:"spawnThingOnRemoved"`
	AddedPartProps        PartProps             `xml:"addedPartProps,omitempty"`
	Stages                Stages                `xml:"stages>li"`
}

type ThingDef struct {
	XMLName               xml.Name              `xml:"ThingDef"`
	DefName               string                `xml:"defName"`
	ParentName            string                `xml:"ParentName,attr"`
	Label                 string                `xml:"label"`
	Description           string                `xml:"description"`
	DescriptionHyperlinks DescriptionHyperlinks `xml:"descriptionHyperlinks"`
	ThingSetMakerTags     string                `xml:"thingSetMakerTags>li"`
	CostList              CostList              `xml:"costList"`
}

type RecipeDef struct {
	XMLName                 xml.Name                 `xml:"RecipeDef"`
	DefName                 string                   `xml:"defName"`
	ParentName              string                   `xml:"ParentName,attr"`
	Label                   string                   `xml:"label"`
	Description             string                   `xml:"description"`
	DescriptionHyperlinks   DescriptionHyperlinks    `xml:"descriptionHyperlinks"`
	JobString               string                   `xml:"jobString"`
	Ingredients             *Ingredients             `xml:"ingredients,omitempty"`
	FixedIngredientFilter   *FixedIngredientFilter   `xml:"fixedIngredientFilter,omitempty"`
	AppliedOnFixedBodyParts *AppliedOnFixedBodyParts `xml:"appliedOnFixedBodyParts,omitempty"`
	AddsHediff              string                   `xml:"addsHediff,omitempty"`
	RemovesHediff           string                   `xml:"removesHediff,omitempty"`
}

type Ingredients struct {
	Ingredients []Ingredient `xml:"li"`
}

type FixedIngredientFilter struct {
	FixedIngredientFilter []string `xml:"thingDefs>li"`
}

type AppliedOnFixedBodyParts struct {
	AppliedOnFixedBodyParts string `xml:"li"`
}

type Defs struct {
	Defs []any
}

type Ingredient struct {
	XMLName xml.Name `xml:"li"`
	Filter  string   `xml:"filter>thingDefs>li"`
	Count   int      `xml:"count"`
}

type DescriptionHyperlinks struct {
	HediffDef []string `xml:"HediffDef,omitempty"`
	RecipeDef []string `xml:"RecipeDef,omitempty"`
	ThingDef  []string `xml:"ThingDef,omitempty"`
}

type PartProps struct {
	Solid bool `xml:"solid,omitempty"`
}

type CostList struct {
	Plasteel        int `xml:"Plasteel,omitempty"`
	ComponentSpacer int `xml:"ComponentSpacer,omitempty"`
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
