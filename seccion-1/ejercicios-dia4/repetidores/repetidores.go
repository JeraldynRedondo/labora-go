/*
- Mostrar el nombre de los planetas del sistema solar, utilizando condiciones, structs y repetidores,
mostrando su cantidad de lunas correctamente.
*/
package main

import (
	"fmt"
)

type Planeta struct {
	nombre string
	NLunas int
}

func main() {

	planetas := []Planeta{
		{"Mercurio", 0},
		{"Venus", 0},
		{"Tierra", 1},
		{"Marte", 2},
		{"Júpiter", 79},
		{"Saturno", 82},
		{"Urano", 27},
		{"Neptuno", 14},
	}

	for i := 0; i < len(planetas); i++ {
		if planetas[i].NLunas == 0 {
			fmt.Printf("El planeta %s No tiene lunas \n", planetas[i].nombre)
		} else if planetas[i].NLunas == 1 {
			fmt.Printf("El planeta %s tiene %d luna \n", planetas[i].nombre, planetas[i].NLunas)

		} else {
			fmt.Printf("El planeta %s tiene %d lunas \n", planetas[i].nombre, planetas[i].NLunas)
		}
	}
}
