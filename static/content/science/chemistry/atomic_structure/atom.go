package atomic_structure

// Atom represents the basic unit of matter
type Atom struct {
	AtomicNumber int
	Element      string
	Protons      int
	Neutrons     int
	Electrons    int
}

// NewAtom creates a new Atom instance. Doesn't take in number of electrons or protons because they should be equal to the atom ic number, if not an ion.
func NewAtom(atomicNumber int, element string, neutrons int) *Atom {
	return &Atom{
		AtomicNumber: atomicNumber,
		Element:      element,
		Protons:      atomicNumber,
		Neutrons:     neutrons,
		Electrons:    atomicNumber,
	}
}

// IsotopeOf checks if two atoms are isotopes of each other
func (a *Atom) IsotopeOf(other *Atom) bool {
	return a.AtomicNumber == other.AtomicNumber && a.Neutrons != other.Neutrons
}

// For more information on electron configuration, see the electron_configuration.go file in this package.
