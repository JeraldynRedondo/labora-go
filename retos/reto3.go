/*
Realice un algoritmo (o sea en un archivo.go codifiquen un programita) para invertir una cadena de caracteres
(la cadena puede ser ingresada por teclado o bien puede ser un valor fijo que ustedes gusten,  BTW a ese valor
fijo dentro del código se lo llama valor "hardcodeado"). ¿Que pasaría se compara una cadena con la misma cadena
al revés? ¿Que problema estoy resolviendo?
*/
package main

import (
	"fmt"
	"strings"
)

func invertirPalabra(hardcodeado string) string {
	Original := []rune(hardcodeado)
	Invertir := make([]rune, len(Original))
	for i, j := 0, len(Original)-1; i < len(Original); i, j = i+1, j-1 {
		Invertir[i] = Original[j]
	}
	return string(Invertir)
}

func main() {
	var hardcodeado string

	fmt.Println("¡Bienvenido! \nPor favor,ingresa la palabra")
	fmt.Scanln(&hardcodeado)

	//hardcodeado = "hola"
	fmt.Printf("La palabra original es: %s \n", hardcodeado)

	hardcodeadoInverso := invertirPalabra(hardcodeado)
	fmt.Printf("La palabra invertida es: %s \n", hardcodeadoInverso)

	minusA := strings.ToLower(hardcodeado)
	minusB := strings.ToLower(hardcodeadoInverso)

	if minusA == minusB {
		fmt.Println("La cadena es un palíndromo")
	} else {
		fmt.Println("La cadena no es un palíndromo")
	}
}
