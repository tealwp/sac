package chemical_bonding

import (
	"fmt"

	"github.com/tealwp/sac/static/content/science/chemistry/atomic_structure"
)

// Bond represents a chemical bond between two atoms
type Bond struct {
	Atom1    *atomic_structure.Atom
	Atom2    *atomic_structure.Atom
	BondType string
}

// FormIonicBond creates an ionic bond between two atoms
func FormIonicBond(metal, nonmetal *atomic_structure.Atom) *Bond {
	return &Bond{
		Atom1:    metal,
		Atom2:    nonmetal,
		BondType: "Ionic",
	}
}

// FormCovalentBond creates a covalent bond between two atoms
func FormCovalentBond(atom1, atom2 *atomic_structure.Atom) *Bond {
	return &Bond{
		Atom1:    atom1,
		Atom2:    atom2,
		BondType: "Covalent",
	}
}

// String returns a string representation of the bond
func (b *Bond) String() string {
	return fmt.Sprintf("%s bond between %s and %s", b.BondType, b.Atom1.Element, b.Atom2.Element)
}

// For more information on molecular geometry, see the geometry.go file in this package.
