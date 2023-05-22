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

	intensity := rand.Intn(6)

	//Variable that calculates the intensity of the attack according to the percentage of life
	calculate := float64(intensity) * float64(p.Life) / float64(p.InitialLife)
	intensityAttack := int(calculate)

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
