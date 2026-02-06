package defMethods

import (
	"encoding/xml"
	"fmt"

	"github.com/Shizu-99/Rimworld-Def-Builder/models"
)

func CreateBodyPart() {
	parts := []string{
		"Eye",
	}
	defList := []any{}
	for _, part := range parts {
		stats := models.StatOffsets{
			ConstructionSpeedFactor: 0.1,
			MiningSpeed:             1.5,
		}
		props := models.PartProps{
			Solid: true,
		}
		caps := models.CapMods{
			Capacities: []models.Capacity{
				{
					Capacity: "Sight",
					Offset:   0.05,
				},
				{
					Capacity: "Moving",
					Offset:   0.5,
				},
			},
		}
		stages := models.Stages{
			Stats:        stats,
			CapacityMods: caps,
		}
		Hediff := models.HediffDef{
			DefName:     part,
			ParentName:  "Makable",
			Label:       part,
			LabelNoun:   "an " + part,
			Description: "An installed" + part + ", its a part of the body.",
			DescriptionHyperlinks: models.DescriptionHyperlinks{
				ThingDef: []string{part},
			},
			SpawnThingOnRemoved: part,
			AddedPartProps:      props,
			Stages:              stages,
		}
		Thing := models.ThingDef{
			DefName:     part,
			ParentName:  "Makable",
			Label:       part,
			Description: "The " + part + ", is a part of the body.",
			DescriptionHyperlinks: models.DescriptionHyperlinks{
				RecipeDef: []string{"Install" + part},
			},
			ThingSetMakerTags: "RewardStandardLowFreq",
			CostList: models.CostList{
				Plasteel:        15,
				ComponentSpacer: 3,
			},
		}
		Install := models.RecipeDef{
			DefName:     "Install" + part,
			ParentName:  "Install",
			Label:       "Install " + part,
			Description: "Install an " + part,
			DescriptionHyperlinks: models.DescriptionHyperlinks{
				HediffDef: []string{part},
				ThingDef:  []string{part},
			},
			JobString: "installing " + part,
			Ingredients: &models.Ingredients{
				Ingredients: []models.Ingredient{
					{
						Filter: part,
						Count:  1,
					},
					{
						Filter: "Medicine",
						Count:  100,
					},
				},
			},
			FixedIngredientFilter: &models.FixedIngredientFilter{
				FixedIngredientFilter: []string{
					part,
					"medicine",
				},
			},
			AppliedOnFixedBodyParts: &models.AppliedOnFixedBodyParts{
				AppliedOnFixedBodyParts: "Eye",
			},
			AddsHediff:    part,
			RemovesHediff: "",
		}
		Remove := models.RecipeDef{
			DefName:     "Remove" + part,
			Label:       "remove " + part,
			Description: "Remove an " + part,
			DescriptionHyperlinks: models.DescriptionHyperlinks{
				ThingDef: []string{
					part,
				},
				HediffDef: []string{
					part,
				},
			},
			Ingredients:   nil,
			JobString:     "Removing " + part,
			RemovesHediff: part,
		}
		defList = append(defList, Hediff, Thing, Install, Remove)
	}
	defs := models.Defs{
		Defs: defList,
	}
	out, _ := xml.MarshalIndent(defs, " ", "	")
	fmt.Println(string(out))
}
