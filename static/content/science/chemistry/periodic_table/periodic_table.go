package periodic_table

import "encoding/json"

// Element represents a single element in the periodic table.
// Specifically, we are talking about a single atom of this element, unless otherwise stated.
//
// Generally speaking the fields we've got on this [Element] are
// known constants -- per element, but may vary slightly from source to source.
// Assume there are edgecases and exceptions, of course.
//
// Fields:
//
// - Symbol e.g., "H", "He", "Fl", "Cl"
//
// - Name e.g., "Hydrogen", "Helium", "Fluorine", "Chlorine"
//
// - AtomicNumber orders the elements on the period table.
//
// - Group is the horizontol placement of the element on the periodic table.
//
// - Period is the vertical placement of the element on the periodic table.
//
// - Category is a general descriptor of the elements properties e.g., "Metal", "Nonmetal", "Metalloid"
//
// - AtomicMass is the mass of a single atom, units are in atomic mass units (amu).
//
// - MeltingPoint is the needed temperature to chemically transform a pure sample of this element from a solid state to a liquid state.
//
// - BoilingPoint is the needed temperature to chemically transform a pure sample of this element from a liquid state to a gaseous state.
//
// - Electronegativity is a measure of how much an atom attracts electrons towards itself when it is bonded to another atom. Important when calculating bonding strength, reactions outcomes and reaction mechanisms.
type Element struct {
	Symbol            string  `json:"symbol"`
	Name              string  `json:"name"`
	AtomicNumber      int     `json:"atomicNumber"`
	Group             int     `json:"group"`
	Period            int     `json:"period"`
	Category          string  `json:"category"`
	AtomicMass        float64 `json:"atomicMass"`
	MeltingPoint      float64 `json:"meltingPoint"`
	BoilingPoint      float64 `json:"boilingPoint"`
	Electronegativity float64 `json:"electronegativity"`
}

// PeriodicTable represents the entire periodic table.
type PeriodicTable struct {
	Elements []*Element `json:"elements"`
}

// NewPeriodicTable creates a new PeriodicTable instance, populating it with some hardcoded elements as an example.
func NewPeriodicTable() *PeriodicTable {
	var pt PeriodicTable
	err := json.Unmarshal([]byte(KnownElements), &pt) // check out the consts file in this directory for some known elements and their chemical properties.
	if err != nil {
		panic(err)
	}
	return &pt
}

// GetElementByName returns the element with the given name.
func (pt *PeriodicTable) GetElementByName(name string) *Element {
	for _, element := range pt.Elements {
		if element.Name == name {
			return element
		}
	}
	return nil
}

// GetElementBySymbol returns the element with the given symbol.
func (pt *PeriodicTable) GetElementBySymbol(symbol string) *Element {
	for _, element := range pt.Elements {
		if element.Symbol == symbol {
			return element
		}
	}
	return nil
}

// GetElementsInGroup returns the elements in the given group.
func (pt *PeriodicTable) GetElementsInGroup(group int) []*Element {
	elements := []*Element{}
	for _, element := range pt.Elements {
		if element.Group == group {
			elements = append(elements, element)
		}
	}
	return elements
}

// GetElementsInPeriod returns the elements in the given period.
func (pt *PeriodicTable) GetElementsInPeriod(period int) []*Element {
	elements := []*Element{}
	for _, element := range pt.Elements {
		if element.Period == period {
			elements = append(elements, element)
		}
	}
	return elements
}

// IsAlkaliMetal checks if the element is an alkali metal.
func (e *Element) IsAlkaliMetal() bool {
	return e.Group == Group1
}

// IsAlkalineEarthMetal checks if the element is an alkaline earth metal.
func (e *Element) IsAlkalineEarthMetal() bool {
	return e.Group == Group2
}

// IsHalogen checks if the element is a halogen.
func (e *Element) IsHalogen() bool {
	return e.Group == Group17
}

// IsNobleGas checks if the element is a noble gas.
func (e *Element) IsNobleGas() bool {
	return e.Group == Group18
}
