package atoms

import (
	"fmt"
	"strings"

	"github.com/tealwp/sac/static/content/science/chemistry/atoms/subatomic"
	"github.com/tealwp/sac/static/content/science/chemistry/atoms/subatomic/electrons"
	"github.com/tealwp/sac/static/content/science/chemistry/atoms/subatomic/neutrons"
	"github.com/tealwp/sac/static/content/science/chemistry/atoms/subatomic/protons"
)

// Atom represents a fundamental unit of matter.
type Atom struct {
	// Element name.
	Element string

	// protons in the nucleus.
	Protons []*protons.Proton

	// neutrons in the nucleus.
	Neutrons []*neutrons.Neutron

	// electrons orbiting the nucleus.
	Electrons []*electrons.Electron
}

// NewAtom creates a new atom with the specified element, proton number, neutron number, and electron number.
func NewAtom(element string, protonNumber, neutronNumber, electronNumber int) *Atom {
	return &Atom{
		Element:   element,
		Protons:   subatomic.NewProtons(protonNumber),
		Neutrons:  subatomic.NewNeutrons(neutronNumber),
		Electrons: subatomic.NewElectrons(electronNumber),
	}
}

// ProtonCount returns the number of protons in the atom.
func (a *Atom) ProtonCount() int {
	return len(a.Protons)
}

// NeutronCount returns the number of neutrons in the atom.
func (a *Atom) NeutronCount() int {
	return len(a.Neutrons)
}

// ElectronCount returns the number of electrons in the atom.
func (a *Atom) ElectronCount() int {
	return len(a.Electrons)
}

// IsNeutral checks if the atom is neutral (equal number of protons and electrons).
func (a *Atom) IsNeutral() bool {
	return a.ProtonCount() == a.ElectronCount()
}

// IsIon checks if the atom is an ion (has a charge).
func (a *Atom) IsIon() bool {
	return !a.IsNeutral()
}

// GetCharge returns the charge of the atom (positive if more protons than electrons, negative if more electrons than protons).
// You'll notice we are just looking at electron and proton count, and the difference between them)
// as opposed to the actualy charge associated with the subatomic particles.
// This is common in high-level chemistry.
func (a *Atom) GetCharge() int {
	return a.ProtonCount() - a.ElectronCount()
}

// CalcCharge returns the calculated charge of the atom.
// This is the more theoretical charge, used in more technical situations.
func (a *Atom) CalcCharge() float64 {
	charge := 0.0
	for _, p := range a.Protons {
		charge += p.Charge
	}
	for _, n := range a.Neutrons {
		charge += n.Charge
	}
	for _, e := range a.Electrons {
		charge += e.Charge
	}
	return charge
}

// GetAtomicMass calculates the approximate atomic mass of the atom based on the sum of protons and neutrons.
// You'll notice we are just looking at neutron and proton count,
// as opposed to the actualy mass associated with the subatomic particles.
// This is common in high-level chemistry.
func (a *Atom) GetAtomicMass() int {
	return a.ProtonCount() + a.NeutronCount()
}

// CalcAtomicMass returns the calculated atomic mass of the atom.
// This is the more theoretical mass, used in more technical situations.
func (a *Atom) CalcAtomicMass() float64 {
	mass := 0.0
	for _, p := range a.Protons {
		mass += p.Mass
	}
	for _, n := range a.Neutrons {
		mass += n.Mass
	}
	for _, e := range a.Electrons {
		mass += e.Mass
	}
	return mass
}

// GetIsotopeName determines the isotope name based on the proton and neutron numbers.
func (a *Atom) GetIsotopeName() string {
	// A simple approach based on the neutron number relative to the proton number
	if a.NeutronCount() == a.ProtonCount() {
		return fmt.Sprintf("%s-%d", a.Element, a.ProtonCount())
	} else if a.NeutronCount() > a.ProtonCount() {
		return fmt.Sprintf("%s-%d", a.Element, a.NeutronCount()+a.ProtonCount())
	} else {
		return fmt.Sprintf("%s-%d", a.Element, a.ProtonCount()-a.NeutronCount())
	}
}

// CheckValenceElectrons determines the number of valence electrons.
func (a *Atom) CheckValenceElectrons() int {
	// A simplified approach based on the electron configuration
	valenceElectrons := 0
	for _, orbital := range strings.Split(electrons.GetOrbitalConfig(a.ElectronCount()), " ") {
		if orbital[0] == '2' {
			valenceElectrons += len(orbital[1:])
		}
	}
	return valenceElectrons
}

func (a *Atom) GetEquationRepresentation() string {
	return ""
}
