package main

import (
	"fmt"
	"time"

	"interfacesGame/fighters"
)

func main() {

	contenders := fighters.CreatedFighters()

	var _, areAtLeastBothAlive = fighters.AreAtLeastBothAlive(contenders)

	for areAtLeastBothAlive {

		for i := 0; i < len(contenders); i++ {
			fighters.Stroke(i, contenders)
		}

		for _, contender := range contenders {
			fmt.Printf("%sLife=%d ", contender.GetName(), contender.GetLife())
		}
		fmt.Println()

		_, areAtLeastBothAlive = fighters.AreAtLeastBothAlive(contenders)

		time.Sleep(3 * time.Second)
	}

	winner, _ := fighters.AreAtLeastBothAlive(contenders)
	fmt.Printf("The winner of the hunger games is %s with a life of %d", winner.GetName(), winner.GetLife())
}
