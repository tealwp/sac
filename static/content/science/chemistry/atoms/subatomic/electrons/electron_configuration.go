package electrons

import "fmt"

// Orbital represents an electron orbital.
type Orbital struct {
	PrincipalQuantumNumber int
	Subshell               string
	Capacity               int
}

// GetOrbitalConfig generates a simplified electron configuration based on the number of electrons.
func GetOrbitalConfig(electrons int) string {
	orbitals := []Orbital{
		{1, "s", 2},
		{2, "s", 2},
		{2, "p", 6},
		{3, "s", 2},
		{3, "p", 6},
		{3, "d", 10},
		{4, "s", 2},
		{4, "p", 6},
		{4, "d", 10},
		{4, "f", 14},
		{5, "s", 2},
		{5, "p", 6},
		{5, "d", 10},
		{5, "f", 14},
		{6, "s", 2},
		{6, "p", 6},
		{6, "d", 10},
		{6, "f", 14},
		{7, "s", 2},
		{7, "p", 6},
		{7, "d", 10},
		{7, "f", 14},
	}

	config := ""
	for _, orbital := range orbitals {
		for electrons >= orbital.Capacity {
			config += fmt.Sprintf("%ds%d ", orbital.PrincipalQuantumNumber, orbital.Capacity)
			electrons -= orbital.Capacity
		}
		if electrons > 0 {
			config += fmt.Sprintf("%ds%d ", orbital.PrincipalQuantumNumber, electrons)
			electrons = 0
		}
	}
	return config
}
