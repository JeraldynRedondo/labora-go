package fighters

import (
	"math/rand"
)

// Criminal it is a structure that represents the criminal fighter, with a base structure.
type Criminal struct {
	BaseFighter
}

// Action to attack. Implementation of the interface <contender> in the structure <Criminal>.
func (p *Criminal) ThrowAttack() int {
	return rand.Intn(7)
}

// Action to recibe attack. Implementation of the interface <contender> in the structure <Criminal>.
func (p *Criminal) RecieveAttack(intensity int) {
	p.Life -= intensity

	if p.Life < 0 {
		p.Life = 0
	}
}

// Name of Fighter <Criminal>.
func (p *Criminal) GetName() string {
	return "Criminal"
}
