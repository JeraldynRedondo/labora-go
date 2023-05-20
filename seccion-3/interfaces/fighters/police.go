package fighters

import (
	"math/rand"
)

// Police it is a structure that represents the police fighter, with a base structure.
type Police struct {
	BaseFighter
	Armour int // 0..50
}

// Action to attack. Implementation of the interface <contender> in the structure <Police>.
func (p Police) ThrowAttack() int {
	return rand.Intn(4)
}

// Action to recibe attack. Implementation of the interface <contender> in the structure <Police>.
func (p *Police) RecieveAttack(intensity int) {
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

// Name of Fighter <Police>.
func (p Police) GetName() string {
	return "Policia"
}
