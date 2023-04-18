/*
Una cadena de ADN se representa como una secuencia circular de bases (adenina, timina, citosina y guanina) que es única

	para cada ser vivo, por ejemplo:

A
T
F
T
C
A
T
G
Dicha cadena se puede representar como un arreglo de caracteres recorriendola en sentido horario desde la parte superior
izquierda:
A T G C G T A T
Se pide diseñar una clase que represente una secuencia de ADN e incluya un método booleano que nos devuelva true si dos
cadenas de ADN coinciden.
MUY IMPORTANTE: La secuencia de ADN es cíclica, por lo que puede comenzar en cualquier posición. Por ejemplo, las dos
secuencias siguientes coinciden:
A T G C G T A T
A T A T G C G T
Pregunta existencial: ¿la cantidad de combinaciones de ADN debe ser finita? ¿Cuántas combinaciones diferentes puede
haber? ¿Puede ocurrir algún día que se empiezan a repetir??Una cadena de ADN se representa como una secuencia circular
de bases (adenina, timina, citosina y guanina) que es única para cada ser vivo, por ejemplo:
*/
package main

import "fmt"

func CompararADN(ADN1, ADN2 string) bool {
	if len(ADN1) != len(ADN2) {
		return false
	}

	for i := 0; i < len(ADN1); i++ {
		for j := 0; j < len(ADN1); j++ {
			if ADN1[i] == ADN2[j] {
				for k := 1; k < len(ADN1); k++ {
					if ADN1[(i+k)%len(ADN1)] != ADN2[(j+k)%len(ADN2)] {
						break
					} else if k == len(ADN1)-1 {
						return true
					}
				}
			}
		}
	}

	return false
}

func main() {
	var ADN1, ADN2 string

	//fmt.Println("¡Bienvenido! \nPor favor,ingresa el primer ADN a comparar")
	//fmt.Scanln(&ADN1)
	//fmt.Println("¡Bienvenido! \nPor favor,ingresa el segundo ADN a comparar")
	//fmt.Scanln(&ADN2)

	//SI
	// ADN1 = "ATGCGTAT"
	// ADN2 = "ATATGCGT"
	//NO
	// ADN1 = "TATGCGJA"
	// ADN2 = "ATATGCGQ"
	//NO
	// ADN1 = "ATGCGTAT"
	// ADN2 = "CGTATAXG"
	//SI
	// ADN1 = "ATGCGTAT"
	// ADN2 = "CGTATATG"
	//SI
	// ADN1 = "ATGCGTAT"
	// ADN2 = "ATATGCGT"

	ADN1 = "123"
	ADN2 = "231"

	resultado := CompararADN(ADN1, ADN2)
	if resultado {
		fmt.Println("Las cadenas coinciden")
	} else {
		fmt.Println("Las cadenas NO coinciden")
	}

}
