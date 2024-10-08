package chemical_reactions

import (
	"fmt"

	"github.com/tealwp/sac/static/content/science/chemistry/bonding"
)

// Reaction represents a chemical reaction
type Reaction struct {
	Reactants []*bonding.Bond
	Products  []*bonding.Bond
	Type      string
}

// NewReaction creates a new chemical reaction
func NewReaction(reactants, products []*bonding.Bond, reactionType string) *Reaction {
	return &Reaction{
		Reactants: reactants,
		Products:  products,
		Type:      reactionType,
	}
}

// BalanceReaction attempts to balance the chemical reaction
func (r *Reaction) BalanceReaction() error {
	// Simplified balancing logic
	// In a real implementation, this would be much more complex
	if len(r.Reactants) == len(r.Products) {
		return nil
	}
	return fmt.Errorf("unable to balance reaction")
}

// String returns a string representation of the reaction
func (r *Reaction) String() string {
	reactants := ""
	for _, bond := range r.Reactants {
		reactants += bond.String() + " + "
	}
	reactants = reactants[:len(reactants)-3] // Remove last " + "

	products := ""
	for _, bond := range r.Products {
		products += bond.String() + " + "
	}
	products = products[:len(products)-3] // Remove last " + "

	return fmt.Sprintf("%s -> %s (%s reaction)", reactants, products, r.Type)
}

// For more information on reaction kinetics, see the kinetics.go file in this package.
