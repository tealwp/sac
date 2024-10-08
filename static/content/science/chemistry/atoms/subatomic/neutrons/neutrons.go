// neutrons is the content space for everything neutron related.
package neutrons

const (
	NeutronMass   = 1.008664915886 // units: atomic mass units (amu)
	NeutronCharge = 0.0            // units: Coulombs
)

// Neutrons are neutral subatomic particles found in the nucleus of an atom.
//
// They have approximately the same mass as protons but do not carry an electric charge.
//
// The number of neutrons in an atom's nucleus, along with the number of protons, determines the atom's mass number.
//
// Neutrons play a crucial role in the stability of the nucleus, helping to counteract the repulsive forces between protons.
//
// Different isotopes of the same element have different numbers of neutrons, but the same number of protons.
//
// Neutrons are involved in nuclear reactions, such as nuclear fission and fusion.
type Neutron struct {
	Charge float64 // Coulombs
	Mass   float64 // atomic mass units (amu)
}

func NewNeutron() *Neutron {
	return &Neutron{
		Charge: NeutronCharge,
		Mass:   NeutronMass,
	}
}
