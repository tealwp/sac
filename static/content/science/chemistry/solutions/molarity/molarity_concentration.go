package molarity_concentration

// Solution represents a chemical solution.
type Solution struct {
    Solute string
    Solvent string
    Molarity float64
    Volume float64 // Liters
    SolventVolume float64 // Liters
}

// CalculateMoles calculates the number of moles of solute in the solution.
func (s *Solution) CalculateMoles() float64 {
    return s.Molarity * s.Volume
}

// CalculateMass calculates the mass of solute in the solution, given the molar mass.
func (s *Solution) CalculateMass(molarMass float64) float64 {
    return s.CalculateMoles() * molarMass
}

// CalculateMassPercentage calculates the mass percentage of the solute in the solution.
func (s *Solution) CalculateMassPercentage(molarMass float64) float64 {
    solutionMass := s.CalculateMass(molarMass) + s.SolventMass(molarMass)
    return (s.CalculateMass(molarMass) / solutionMass) * 100
}

// CalculateMolality calculates the molality of the solution, given the density.
func (s *Solution) CalculateMolality(density float64) float64 {
    solutionMass := density * s.Volume * 1000 // Convert volume to mL
    return s.CalculateMoles() / (solutionMass / 1000) // Convert mass to kg
}

// CalculateVolumePercentage calculates the volume percentage of the solute in the solution.
func (s *Solution) CalculateVolumePercentage() float64 {
    return (s.Volume / (s.Volume + s.SolventVolume)) * 100
}

// CalculatePartsPerMillion calculates the parts per million (ppm) of the solute in the solution.
func (s *Solution) CalculatePartsPerMillion(molarMass float64) float64 {
    solutionMass := s.CalculateMass(molarMass) + s.SolventMass(molarMass)
    return (s.CalculateMass(molarMass) / solutionMass) * (10^6)
}

func (s *Solution) SolventMass(molarMass float64) float64 {
    // Assuming the total mass of the solution is known
    totalMass := 100.0 // Replace with the actual total mass in grams

    // Calculate the mass of the solute
    soluteMass := s.CalculateMass(molarMass) // Replace molarMass with the actual molar mass

    return totalMass - soluteMass
}
