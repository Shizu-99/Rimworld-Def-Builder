package defMethods

import (
	"encoding/xml"
	"fmt"

	"github.com/Shizu-99/Rimworld-Def-Builder/models"
)

func CreateBodyPart() {
	parts := []string{
		"Eye",
		"Leg",
		"Arm",
	}
	defList := []any{}
	for _, part := range parts {
		stats := &models.StatOffsets{
			ConstructionSpeedFactor: 0.1,
			MiningSpeed:             1.5,
		}
		props := &models.PartProps{
			Solid: true,
		}
		caps := &models.CapMods{
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
		stages := &models.Stages{
			Stats:        *stats,
			CapacityMods: *caps,
		}
		Hediff := &models.HediffDef{
			DefName:               part,
			ParentName:            "Makable",
			Label:                 part,
			LabelNoun:             "an " + part,
			Description:           "An " + part + ", its a part of the body.",
			DescriptionHyperlinks: part,
			SpawnThingOnRemoved:   part,
			AddedPartProps:        *props,
			Stages:                *stages,
		}
		Thing := &models.ThingDef{
			DefName:    part,
			ParentName: "Makable",
		}
		Install := &models.RecipeDef{
			DefName:    part,
			ParentName: "Install",
		}
		Remove := &models.RecipeDef{
			DefName:    part,
			ParentName: "Remove",
		}
		defList = append(defList, Hediff, Thing, Install, Remove)
	}
	defs := models.Defs{
		Defs: defList,
	}
	out, _ := xml.MarshalIndent(defs, " ", "	")
	fmt.Println(string(out))
}
