// electrons is the content space for everything electron related.
package electrons

const (
	ElectronMass   = 0.00054858 // units: atomic mass units (amu)
	ElectronCharge = -1.602e-19 // units: Coulombs
)

// Electrons are the negatively charged, nearly weightless, subatomic particle that make up an atom.
//
// They can be found in clouds at different levels and distances around the nucleus.
// More often than not these clouds are represented as orbits, and for most purposes,
// can be treated as orbitals like the moon, following a known path in space. This simplifies the needed
// knowledge to understand most concepts regarding electrons and atoms.
//
// As far as their importance in the chemical make-up of an atom, some think they are the most important factor.
//
// Electrons are shared or transferred between atoms when bonding. This process depends on the bond type.
//
// An electron configuration is our way of defining the layout of an atom's electrons into these layers.
// This configuration is a great way to understand if the atom is susceptible to bonding, or extremely hard to bond with.
type Electron struct {
	Charge float64 // Coulombs
	Mass   float64 // atomic mass units (amu)
}

// NewElectron creates a new electron.
func NewElectron() *Electron {
	return &Electron{
		Charge: ElectronCharge,
		Mass:   ElectronMass,
	}
}
