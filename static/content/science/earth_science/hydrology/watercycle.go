package watercycle

import "fmt"

// WaterParticle represents a particle of water in the water cycle
type WaterParticle struct {
	State    string // Can be "liquid", "solid", or "gas"
	Location string
}

// Evaporate changes the state of water from liquid to gas
func (w *WaterParticle) Evaporate() {
	if w.State == "liquid" {
		w.State = "gas"
		w.Location = "atmosphere"
		fmt.Println("Water evaporated into the atmosphere")
	}
}

// Condense changes the state of water from gas to liquid
func (w *WaterParticle) Condense() {
	if w.State == "gas" {
		w.State = "liquid"
		w.Location = "cloud"
		fmt.Println("Water vapor condensed into a cloud")
	}
}

// Precipitate moves water from a cloud to the ground
func (w *WaterParticle) Precipitate() {
	if w.State == "liquid" && w.Location == "cloud" {
		w.Location = "ground"
		fmt.Println("Water precipitated to the ground")
	}
}

// SimulateWaterCycle runs a simple simulation of the water cycle
func SimulateWaterCycle() {
	water := WaterParticle{State: "liquid", Location: "ocean"}

	fmt.Println("Starting water cycle simulation:")
	fmt.Printf("Initial state: %+v\n", water)

	water.Evaporate()
	water.Condense()
	water.Precipitate()

	fmt.Printf("Final state: %+v\n", water)
}
