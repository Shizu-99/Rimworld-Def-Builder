package main

import (
	"encoding/xml"
	"fmt"

	"github.com/Shizu-99/Rimworld-Def-Builder/models"
)

func main() {
	parts := []string{
		"Eye",
		"Leg",
		"Arm",
	}
	defList := []any{}
	for _, part := range parts {
		Hediff := &models.HediffDef{
			DefName:    part,
			ParentName: "Makable",
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
	out, _ := xml.MarshalIndent(defs, " ", " ")
	fmt.Println(string(out))
}
