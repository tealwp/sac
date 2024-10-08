// protons is the content space for everything proton related.
package protons

const (
	ProtonMass   = 1.0073    // units: atomic mass units (amu)
	ProtonCharge = 1.602e-19 // units: Coulombs
)

// Protons are positively charged subatomic particles found in the nucleus of an atom.
//
// Their number determines the atomic number of an element.
// Generally speaking an atom will have the same number of neutrons as it does protons.
// Protons and neutrons are held together in the nucleus by the strong nuclear force.
//
// The number of protons in an atom's nucleus determines its chemical properties,
// as it affects the number and arrangement of electrons.
//
// Protons play a crucial role in the stability and structure of atoms.
type Proton struct {
	Charge float64 // Coulombs
	Mass   float64 // atomic mass units (amu)
}

// NewProton creates a new proton.
func NewProton() *Proton {
	return &Proton{
		Charge: ProtonCharge,
		Mass:   ProtonMass,
	}
}
