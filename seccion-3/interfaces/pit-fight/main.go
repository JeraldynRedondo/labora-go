package main

import (
	"fmt"
	"time"

	"interfacesGame/repository"
)

func main() {

	contenders := repository.CreatedFighters()

	var areAtLeastBothAlive = repository.CountFightersAlive(contenders)

	for areAtLeastBothAlive {

		for i := 0; i < len(contenders); i++ {
			repository.Stroke(i, contenders)
		}

		for _, contender := range contenders {
			fmt.Printf("%sLife=%d ", contender.GetName(), contender.GetLife())
		}
		fmt.Println()

		areAtLeastBothAlive = repository.CountFightersAlive(contenders)

		time.Sleep(3 * time.Second)
	}
}
