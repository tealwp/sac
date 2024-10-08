package periodic_table

const (
	// KnownElements are the first handful of elements with known chemical property values.
	KnownElements = `
{
	"elements": [
		{ "symbol": "H", "name": "Hydrogen", "atomicNumber": 1, "group": 1, "period": 1, "category": "Nonmetal", "atomicMass": 1.008, "meltingPoint": -259.14, "boilingPoint": -252.87, "electronegativity": 2.2 },
		{ "symbol": "He", "name": "Helium", "atomicNumber": 2, "group": 18, "period": 1, "category": "Noble Gas", "atomicMass": 4.003, "meltingPoint": -272.2, "boilingPoint": -268.9, "electronegativity": 0.9 },
		{ "symbol": "Li", "name": "Lithium", "atomicNumber": 3, "group": 1, "period": 2, "category": "Alkali Metal", "atomicMass": 6.941, "meltingPoint": 180.54, "boilingPoint": 1342, "electronegativity": 0.98 },
		{ "symbol": "Be", "name": "Beryllium", "atomicNumber": 4, "group": 2, "period": 2, "category": "Alkaline Earth Metal", "atomicMass": 9.012, "meltingPoint": 1287, "boilingPoint": 2970, "electronegativity": 1.57 },
		{ "symbol": "B", "name": "Boron", "atomicNumber": 5, "group": 13, "period": 2, "category": "Metalloid", "atomicMass": 10.811, "meltingPoint": 2349, "boilingPoint": 4200, "electronegativity": 2.04 },
		{ "symbol": "C", "name": "Carbon", "atomicNumber": 6, "group": 14, "period": 2, "category": "Nonmetal", "atomicMass": 12.011, "meltingPoint": 3550, "boilingPoint": 4827, "electronegativity": 2.55 },
		{ "symbol": "N", "name": "Nitrogen", "atomicNumber": 7, "group": 15, "period": 2, "category": "Nonmetal", "atomicMass": 14.007, "meltingPoint": -210, "boilingPoint": -195.8, "electronegativity": 3.04 },
		{ "symbol": "O", "name": "Oxygen", "atomicNumber": 8, "group": 16, "period": 2, "category": "Nonmetal", "atomicMass": 15.999, "meltingPoint": -218.4, "boilingPoint": -183.0, "electronegativity": 3.44 },
		{ "symbol": "F", "name": "Fluorine", "atomicNumber": 9, "group": 17, "period": 2, "category": "Halogen", "atomicMass": 18.998, "meltingPoint": -223, "boilingPoint": -188.1, "electronegativity": 3.98 },
		{ "symbol": "Ne", "name": "Neon", "atomicNumber": 10, "group": 18, "period": 2, "category": "Noble Gas", "atomicMass": 20.179, "meltingPoint": -248.6, "boilingPoint": -246.1, "electronegativity": 0.9 },
		{ "symbol": "Na", "name": "Sodium", "atomicNumber": 11, "group": 1, "period": 3, "category": "Alkali Metal", "atomicMass": 22.990, "meltingPoint": 97.8, "boilingPoint": 883, "electronegativity": 0.93 },
		{ "symbol": "Mg", "name": "Magnesium", "atomicNumber": 12, "group": 2, "period": 3, "category": "Alkaline Earth Metal", "atomicMass": 24.305, "meltingPoint": 650, "boilingPoint": 1090, "electronegativity": 1.31 },
		{ "symbol": "Al", "name": "Aluminum", "atomicNumber": 13, "group": 13, "period": 3, "category": "Post-Transition Metal", "atomicMass": 26.982, "meltingPoint": 660.3, "boilingPoint": 2467, "electronegativity": 1.61 },
		{ "symbol": "Si", "name": "Silicon", "atomicNumber": 14, "group": 14, "period": 3, "category": "Metalloid", "atomicMass": 28.086, "meltingPoint": 1414, "boilingPoint": 3265, "electronegativity": 1.90 },
		{ "symbol": "P", "name": "Phosphorus", "atomicNumber": 15, "group": 15, "period": 3, "category": "Nonmetal", "atomicMass": 30.974, "meltingPoint": 44.15, "boilingPoint": 280, "electronegativity": 2.19 },
		{ "symbol": "S", "name": "Sulfur", "atomicNumber": 16, "group": 16, "period": 3, "category": "Nonmetal", "atomicMass": 32.065, "meltingPoint": 115.2, "boilingPoint": 444.6, "electronegativity": 2.58 },
		{ "symbol": "Cl", "name": "Chlorine", "atomicNumber": 17, "group": 17, "period": 3, "category": "Halogen", "atomicMass": 35.453, "meltingPoint": -101.5, "boilingPoint": -34.04, "electronegativity": 3.16 },
		{ "symbol": "Ar", "name": "Argon", "atomicNumber": 18, "group": 18, "period": 3, "category": "Noble Gas", "atomicMass": 39.948, "meltingPoint": -189.3, "boilingPoint": -185.7, "electronegativity": 0.9 },
		{ "symbol": "K", "name": "Potassium", "atomicNumber": 19, "group": 1, "period": 4, "category": "Alkali Metal", "atomicMass": 39.098, "meltingPoint": 63.5, "boilingPoint": 761, "electronegativity": 0.82 },
		{ "symbol": "Ca", "name": "Calcium", "atomicNumber": 20, "group": 2, "period": 4, "category": "Alkaline Earth Metal", "atomicMass": 40.078, "meltingPoint": 842, "boilingPoint": 1484, "electronegativity": 1.00 },
		{ "symbol": "Sc", "name": "Scandium", "atomicNumber": 21, "group": 3, "period": 4, "category": "Transition Metal", "atomicMass": 44.956, "meltingPoint": 1541, "boilingPoint": 2836, "electronegativity": 1.36 },
		{ "symbol": "Ti", "name": "Titanium", "atomicNumber": 22, "group": 4, "period": 4, "category": "Transition Metal", "atomicMass": 47.867, "meltingPoint": 1668, "boilingPoint": 3287, "electronegativity": 1.54 },
		{ "symbol": "V", "name": "Vanadium", "atomicNumber": 23, "group": 5, "period": 4, "category": "Transition Metal", "atomicMass": 50.942, "meltingPoint": 1910, "boilingPoint": 3407, "electronegativity": 1.63 }
	]
}
`
)

// Constants representing common periodic table groups.
// There are more, but these are some of the most strongly defined groups.
// Elements in groups share a lot of properties.
const (
	Group1  = 1  // Alkali Metals
	Group2  = 2  // Alkaline Earth Metals
	Group17 = 17 // Halogens
	Group18 = 18 // Noble Gases
)
