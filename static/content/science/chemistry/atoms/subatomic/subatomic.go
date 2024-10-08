// subatomic is representing the building blocks of an atom.
//
// You thought an atom was small? Subatomic particles make up an atom.
//
// They directly responsible for how it interacts with other atoms
// during different processes like bonding or reacting.
//
// In chemistry (and physics) there are dozens of subatomic particles.
// Protons, neutrons and electrons are the three you will be using in the `science/chemistry` content space,
// as the other known subatomic particles are much more complex and often theoretical. These three typically
// suffice when trying to understand an atoms properties.
package subatomic

import (
	"github.com/tealwp/sac/static/content/science/chemistry/atoms/subatomic/electrons"
	"github.com/tealwp/sac/static/content/science/chemistry/atoms/subatomic/neutrons"
	"github.com/tealwp/sac/static/content/science/chemistry/atoms/subatomic/protons"
)

// NewProtons creates a slice of [*protons.Proton], with size as an input
func NewProtons(n int) []*protons.Proton {
	p := make([]*protons.Proton, n)
	for i := 0; i < n; i++ {
		p[i] = protons.NewProton()
	}
	return p
}

// NewElectrons creates a slice of [*electrons.Electron], with size as an input
func NewElectrons(n int) []*electrons.Electron {
	e := make([]*electrons.Electron, n)
	for i := 0; i < n; i++ {
		e[i] = electrons.NewElectron()
	}
	return e
}

// NewNeutrons creates a slice of [*neutrons.Neutron], with size as an input
func NewNeutrons(n int) []*neutrons.Neutron {
	ne := make([]*neutrons.Neutron, n)
	for i := 0; i < n; i++ {
		ne[i] = neutrons.NewNeutron()
	}
	return ne
}
