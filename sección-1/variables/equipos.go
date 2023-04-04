/*
Crea un archivo llamado equipos.go que muestre en su ejecución los nombres de 5 equipos de fútbol junto con la cantidad
de títulos que han ganado, uno debajo del otro.
*/
package main

import (
	"fmt"
)

func main() {

	var titulos1, titulos2, titulos3, titulos4, titulos5 = 14, 6, 5, 3, 2

	var (
		equipo1 = "1,Real Madrid C.F."
		equipo2 = "2, F.C. Bayern Múnich"
		equipo3 = "3, F C. Barcelona"
		equipo4 = "4,Manchester United"
		equipo5 = "5, Juventus F.C."
	)

	fmt.Printf("El equipo %s ha ganado  %v títulos de La UEFA Champions League \n", equipo1, titulos1)
	fmt.Printf("El equipo %s ha ganado  %v títulos de La UEFA Champions League \n", equipo2, titulos2)
	fmt.Printf("El equipo %s ha ganado  %v títulos de La UEFA Champions League \n", equipo3, titulos3)
	fmt.Printf("El equipo %s ha ganado  %v títulos de La UEFA Champions League \n", equipo4, titulos4)
	fmt.Printf("El equipo %s ha ganado  %v títulos de La UEFA Champions League \n", equipo5, titulos5)
}
