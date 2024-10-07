package acid_base

import (
	"math"
)

type Solution struct {
    pH float64
}

func CalculatepH(hydrogenIonConcentration float64) float64 {
    return -math.Log10(hydrogenIonConcentration)
}

func CalculateHydrogenIonConcentration(pH float64) float64 {
    return math.Pow(10, -pH)
}

// Represents a strong acid.
type StrongAcid struct {
    Concentration float64
}

// Represents a strong base.
type StrongBase struct {
    Concentration float64
}

// Represents a weak acid.
type WeakAcid struct {
    Concentration float64
    Ka float64
}

// Represents a weak base.
type WeakBase struct {
    Concentration float64
    Kb float64
}

// Neutralizes a strong acid with a strong base.
func NeutralizeStrongAcid(acid StrongAcid, base StrongBase) float64 {
    // Calculate the moles of acid and base
    molesAcid := acid.Concentration
    molesBase := base.Concentration

    // Determine the limiting reactant
    if molesAcid > molesBase {
        // Acid is in excess
        return CalculatepH(molesAcid - molesBase)
    } else if molesAcid < molesBase {
        // Base is in excess
        return 14 - CalculatepH(molesBase - molesAcid)
    } else {
        // Neutralization
        return 7.0
    }
}

// Neutralizes a weak acid with a strong base.
func NeutralizeWeakAcid(acid WeakAcid, base StrongBase) float64 {
    // Calculate the moles of acid and base
    molesAcid := acid.Concentration
    molesBase := base.Concentration

    // Determine the limiting reactant
    if molesAcid > molesBase {
        // Acid is in excess
        return CalculatepH(molesAcid - molesBase)
    } else if molesAcid < molesBase {
        // Base is in excess
        return 14 - CalculatepH(molesBase - molesAcid)
    } else {
        // Neutralization
        // Calculate the pH using the Henderson-Hasselbalch equation
        return acid.Ka + math.Log10(molesBase / molesAcid)
    }
}

// ... (add functions for other types of acid-base reactions, such as weak acid-weak base neutralization)