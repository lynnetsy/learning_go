package main

import (
	"fmt"
)

func main() {
	// ns es la variable que va a obtener la funcion
	ns := GetNutritionalScore(NutritionalData{
		Energy:              EnergyFromKcal(),
		Sugars:              SugarGram(),
		SaturatedFattyAcids: SaturatedFattyAcids(),
		Sodium:              SodiumMiligram(),
		Fruits:              FuitsPercent(),
		Fibre:               FiberGram(),
		Protein:             ProteinGram(),
	}, Food)

	fmt.Printf("Nutritional Score: %d", ns.Value)
	fmt.Printf("NutriScore: %s\n", ns.GetNutriScore())
}
