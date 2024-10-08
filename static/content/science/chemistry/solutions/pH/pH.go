package ph

import (
	"math"
)

// Compound represents a chemical compound in a solution.
type Compound struct {
	Name                     string
	Concentration            float64
	Volume                   float64
	HydrogenIonConcentration float64
}

// Solution represents a solution containing two compounds.
type Solution struct {
	Compound1 Compound
	Compound2 Compound
}

// CalculatePH calculates the pH from a given hydrogen ion concentration.
func CalculatePH(hydrogenIonConcentration float64) float64 {
	return -math.Log10(hydrogenIonConcentration)
}

// CalculateHydrogenIonConcentration calculates the hydrogen ion concentration from a given pH.
func CalculateHydrogenIonConcentration(pH float64) float64 {
	return math.Pow(10.0, -pH)
}

// IsAcidic determines if a solution is acidic based on its pH.
func IsAcidic(pH float64) bool {
	return pH < 7
}

// IsBasic determines if a solution is basic based on its pH.
// Also known as alkaline.
func IsBasic(pH float64) bool {
	return pH > 7
}

// IsNeutral determines if a solution is neutral based on its pH.
func IsNeutral(pH float64) bool {
	return pH == 7
}

// CalculatePHDifference calculates the difference in pH between two solutions.
func CalculatePHDifference(pH1, pH2 float64) float64 {
	return math.Abs(float64(pH1) - float64(pH2))
}

// CalculateConcentrationRatio calculates the ratio of hydrogen ion concentrations between two solutions.
func CalculateConcentrationRatio(pH1, pH2 float64) float64 {
	return math.Pow(10.0, float64(pH2)-float64(pH1))
}

// CalculateSolutionPH calculates the pH of the solution based on the pH, concentration, and volume of its compounds.
func (s Solution) CalculateSolutionPH() float64 {
	// Calculate the total moles of hydrogen ions from each compound.
	totalMolesHydrogenIon1 := s.Compound1.HydrogenIonConcentration * s.Compound1.Volume
	totalMolesHydrogenIon2 := s.Compound2.HydrogenIonConcentration * s.Compound2.Volume

	// Calculate the total volume of the solution.
	totalVolume := s.Compound1.Volume + s.Compound2.Volume

	// Calculate the total concentration of hydrogen ions in the solution.
	totalHydrogenIonConcentration := (totalMolesHydrogenIon1 + totalMolesHydrogenIon2) / totalVolume

	// Calculate the pH of the solution.
	solutionPH := -math.Log10(float64(totalHydrogenIonConcentration))

	return solutionPH
}
