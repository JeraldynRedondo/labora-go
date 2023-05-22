package fighters

import (
	"fmt"
	"math/rand"
	"time"
)

// CreatedFighters it is a function that returns a vector of <Fighters> with the implementation of the interface <Contender>.
func CreatedFighters() []Contender {
	numberOfFighters := 3
	var Contenders []Contender = make([]Contender, numberOfFighters)

	rand.Seed(time.Now().UnixNano())
	randomValue := rand.Intn(numberOfFighters)

	Contenders[randomValue] = &police
	Contenders[(randomValue+1)%numberOfFighters] = &criminal
	Contenders[(randomValue+2)%numberOfFighters] = &paladin

	for i, contender := range Contenders {
		fmt.Printf("Contender %d: %s Life= %d\n", i, contender.GetName(), contender.GetLife())
	}

	return Contenders
}

// AreAtLeastBothAlive it is a function that returns if are at least two fighters alive and the winner.
func AreAtLeastBothAlive(contenders []Contender) (Contender, bool) {
	count := 0
	var possibleWinner Contender
	for _, contender := range contenders {
		if contender.IsAlive() {
			count++
			possibleWinner = contender
		}
	}

	if count > 1 {
		return nil, true
	}

	return possibleWinner, false
}

// Stroke it is a function that executes the attack of a fighter to the other fighters
func Stroke(fighter int, contenders []Contender) {
	if contenders[fighter].IsAlive() {
		for i := 0; i < len(contenders); i++ {
			if i != fighter && contenders[i].IsAlive() {
				intensity := contenders[fighter].ThrowAttack()
				fmt.Println(contenders[fighter].GetName(), " throw punch with intensity =", intensity, "to ", contenders[i].GetName())
				contenders[i].RecieveAttack(intensity)
			}
		}
	}
}
