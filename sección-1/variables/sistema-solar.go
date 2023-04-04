/*
Crear un archivo llamado sistema-solar.go que muestre en su ejecución, todos los planetas del sistema solar junto
con su numero de lunas, uno debajo del otro.
*/
package main

import (
	"fmt"
)

func main() {

	var nLunas1, nLunas2, nLunas3, nLunas4, nLunas5, nLunas6, nLunas7, nLunas8 = 0, 0, 1, 2, 79, 82, 27, 14

	var (
		planeta1 = "Mercurio"
		planeta2 = "Venus"
		planeta3 = "Tierra"
		planeta4 = "Marte"
		planeta5 = "Júpiter"
		planeta6 = "Saturno"
		planeta7 = "Urano"
		planeta8 = "Neptuno"
	)

	fmt.Printf("El planeta %s tiene %d lunas \n", planeta1, nLunas1)
	fmt.Printf("El planeta %s tiene %d lunas \n", planeta2, nLunas2)
	fmt.Printf("El planeta %s tiene %d lunas \n", planeta3, nLunas3)
	fmt.Printf("El planeta %s tiene %d lunas \n", planeta4, nLunas4)
	fmt.Printf("El planeta %s tiene %d lunas \n", planeta5, nLunas5)
	fmt.Printf("El planeta %s tiene %d lunas \n", planeta6, nLunas6)
	fmt.Printf("El planeta %s tiene %d lunas \n", planeta7, nLunas7)
	fmt.Printf("El planeta %s tiene %d lunas \n", planeta8, nLunas8)
}
