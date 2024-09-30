package atomic_structure

import "fmt"

// ElectronConfiguration represents the arrangement of electrons in an atom
type ElectronConfiguration struct {
	Atom   *Atom
	Shells []int
}

// DetermineElectronConfiguration calculates the electron configuration for an atom
func DetermineElectronConfiguration(atom *Atom) *ElectronConfiguration {
	// Simplified electron configuration calculation
	shells := []int{2, 8, 18, 32}
	config := &ElectronConfiguration{Atom: atom}

	remaining := atom.Electrons
	for _, shellCapacity := range shells {
		if remaining > shellCapacity {
			config.Shells = append(config.Shells, shellCapacity)
			remaining -= shellCapacity
		} else {
			config.Shells = append(config.Shells, remaining)
			break
		}
	}

	return config
}

// String returns a string representation of the electron configuration
func (ec *ElectronConfiguration) String() string {
	return fmt.Sprintf("Electron Configuration for %s: %v", ec.Atom.Element, ec.Shells)
}

// For more information on chemical bonding, see the bonding.go file in the chemical_bonding package.
