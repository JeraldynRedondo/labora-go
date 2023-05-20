package repository

import (
	"fmt"
	"interfacesGame/fighters"
	"interfacesGame/model"
	"math/rand"
	"time"
)

// CreatedFighters it is a function that returns a vector of <Fighters> with the implementation of the interface <Contender>.
func CreatedFighters() []fighters.Contender {
	numberOfFighters := 3
	var Contenders []fighters.Contender = make([]fighters.Contender, numberOfFighters)

	rand.Seed(time.Now().UnixNano())
	randomValue := rand.Intn(numberOfFighters)

	Contenders[randomValue] = &model.Police
	Contenders[(randomValue+1)%numberOfFighters] = &model.Criminal
	Contenders[(randomValue+2)%numberOfFighters] = &model.Paladin

	for i, contender := range Contenders {
		fmt.Printf("Contender %d: %s Life= %d\n", i, contender.GetName(), contender.GetLife())
	}

	return Contenders
}

// CountFightersAlive it is a function that returns the number of fighters alive.
func CountFightersAlive(contenders []fighters.Contender) bool {
	count := 0
	for _, contender := range contenders {
		if contender.IsAlive() {
			count++
		}
	}

	if count > 1 {
		return true
	}

	return false
}

// Stroke it is a function that executes the attack of a fighter to the other fighters
func Stroke(fighter int, contenders []fighters.Contender) {
	if contenders[fighter].IsAlive() {
		intensity := contenders[fighter].ThrowAttack()
		fmt.Println(contenders[fighter].GetName(), " tira golpe con intensidad =", intensity)

		for i := 0; i < len(contenders); i++ {
			if i != fighter {
				contenders[i].RecieveAttack(intensity)
			}
		}

	}
}
