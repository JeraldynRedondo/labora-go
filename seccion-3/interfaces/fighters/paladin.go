package fighters

import (
	"math/rand"
)

// Paladin it is a structure that represents the paladin fighter, with a base structure.
type Paladin struct {
	BaseFighter
	Armour int // 0..50
}

// Action to attack. Implementation of the interface <contender> in the structure <Paladin>.
func (p Paladin) ThrowAttack() int {
	intensityAttack := rand.Intn(6)

	if p.Life <= (p.InitialLife / 2) {
		return intensityAttack / 2
	}

	return intensityAttack
}

// Action to recibe attack. Implementation of the interface <contender> in the structure <Paladin>.
func (p *Paladin) RecieveAttack(intensity int) {
	if p.Armour > 0 {
		diff := (p.Armour - intensity)
		hasArmour := diff > 0
		if hasArmour {
			p.Armour -= intensity
			intensity = 0
		} else {
			p.Armour = 0
			intensity = -diff // intensity -= p.Armour
		}
	}

	p.Life -= intensity

	if p.Life < 0 {
		p.Life = 0
	}
}

// Name of Fighter <Paladin>.
func (p Paladin) GetName() string {
	return "Paladin"
}
