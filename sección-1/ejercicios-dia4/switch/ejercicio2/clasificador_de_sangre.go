/*
Escribe un programa en Go que clasifique a las personas según su tipo de sangre. Crea una función que tome como entrada
el tipo de sangre como cadena y devuelva una cadena con la descripción de la clasificación. Utiliza la declaración switch
para determinar la clasificación según el tipo de sangre ingresado. Por ejemplo, si se ingresa "AB+", el programa debe
devolver "Grupo sanguíneo AB, factor Rh positivo".
*/
package main

import (
	"fmt"
)

func main() {
	var TipodeSangre string
	fmt.Println("¡Bienvenido! \nPor favor,ingresa tu tipo de sangre sin separaciones y en mayúscula. Si desea terminar el problema ingrese 1.")
	fmt.Scanf("%s", &TipodeSangre)
	clasificarSangre(TipodeSangre)

	//clasificarSangre("o-")

}

func clasificarSangre(TipodeSangre string) {
	switch TipodeSangre {
	case "A+":
		fmt.Println("Grupo sanguíneo A, factor Rh positivo")
	case "O+":
		fmt.Println("Grupo sanguíneo O, factor Rh positivo")
	case "B+":
		fmt.Println("Grupo sanguíneo B, factor Rh positivo")
	case "AB+":
		fmt.Println("Grupo sanguíneo AB, factor Rh positivo")
	case "A-":
		fmt.Println("Grupo sanguíneo A, factor Rh negativo")
	case "O-":
		fmt.Println("Grupo sanguíneo O, factor Rh negativo")
	case "B-":
		fmt.Println("Grupo sanguíneo B, factor Rh negativo")
	case "AB-":
		fmt.Println("Grupo sanguíneo AB, factor Rh negativo")
	default:
		fmt.Printf("¡Entrada incorrecta, el tipo de sangre %s, no corresponde a ningún grupo sanguíneo o factor RH!\nPor favor no utilice espacios, escriba en mayúscula, verifique la información suministrada e intente de nuevo\n", TipodeSangre)
	}
}
