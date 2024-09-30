package chemical_bonding

import (
	"fmt"

	"github.com/tealwp/sac/static/content/science/chemistry/atomic_structure"
)

// MolecularGeometry represents the three-dimensional arrangement of atoms in a molecule
type MolecularGeometry struct {
	CentralAtom *atomic_structure.Atom
	Shape       string
	BondAngle   float64
}

// PredictGeometry predicts the molecular geometry based on the number of electron domains
func PredictGeometry(centralAtom *atomic_structure.Atom, electronDomains int) *MolecularGeometry {
	var shape string
	var bondAngle float64

	switch electronDomains {
	case 2:
		shape = "Linear"
		bondAngle = 180.0
	case 3:
		shape = "Trigonal Planar"
		bondAngle = 120.0
	case 4:
		shape = "Tetrahedral"
		bondAngle = 109.5
	case 5:
		shape = "Trigonal Bipyramidal"
		bondAngle = 90.0 // Simplified, actual angles vary
	case 6:
		shape = "Octahedral"
		bondAngle = 90.0
	default:
		shape = "Unknown"
		bondAngle = 0.0
	}

	return &MolecularGeometry{
		CentralAtom: centralAtom,
		Shape:       shape,
		BondAngle:   bondAngle,
	}
}

// String returns a string representation of the molecular geometry
func (mg *MolecularGeometry) String() string {
	return fmt.Sprintf("Molecular Geometry: %s, Bond Angle: %.1f°", mg.Shape, mg.BondAngle)
}

// For more information on chemical reactions, see the reactions.go file in the chemical_reactions package.
