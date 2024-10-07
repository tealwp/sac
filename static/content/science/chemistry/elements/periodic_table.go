package elements

// Constants representing common periodic table groups
const (
    Group1 = 1 // Alkali Metals
    Group2 = 2 // Alkaline Earth Metals
    Group17 = 17 // Halogens
    Group18 = 18 // Noble Gases
)

// Element represents a single element in the periodic table.
type Element struct {
    Symbol     string
    Name       string
    AtomicNumber int
    Group      int
    Period     int
    Category   string // e.g., "Metal", "Nonmetal", "Metalloid"
    AtomicMass float64
    ElectronConfiguration string
    MeltingPoint float64
    BoilingPoint float64
    Electronegativity float64
}

// PeriodicTable represents the entire periodic table.
type PeriodicTable struct {
    Elements []*Element
}

// NewPeriodicTable creates a new PeriodicTable instance, populating it with some hardcoded elements as an example.
func NewPeriodicTable() *PeriodicTable {
    elements := []*Element{
        {
            Symbol:     "H",
            Name:       "Hydrogen",
            AtomicNumber: 1,
            Group:      1,
            Period:     1,
            Category:   "Nonmetal",
            AtomicMass: 1.008,
            ElectronConfiguration: "1s1",
        },
        {
            Symbol:     "He",
            Name:       "Helium",
            AtomicNumber: 2,
            Group:      18,
            Period:     1,
            Category:   "Noble Gas",
            AtomicMass: 4.003,
            ElectronConfiguration: "1s2",
        },
        // Add more elements here
    }

    periodicTable := &PeriodicTable{
        Elements: elements,
    }

    return periodicTable
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

// IsHalogen checks if the element is a halogen.
func (e *Element) IsHalogen() bool {
    return e.Group == Group17
}

// IsNobleGas checks if the element is a noble gas.
func (e *Element) IsNobleGas() bool {
    return e.Group == Group18
}