/*
Crea un archivo llamado restaurantes.go que muestre en su ejecución los nombres y las calificaciones de 3 restaurantes
diferentes, uno debajo del otro.
*/
package main

import (
	"fmt"
)

func main() {

	var calificacion1, calificacion2, calificacion3 float32 = 4.5, 4.0, 5.0

	var (
		restaurante1 = "1,Las delicias"
		restaurante2 = "2, Sabor de mi tierra"
		restaurante3 = "3, KFC"
	)

	fmt.Printf("El restaurante %s tiene  una calificación de %v \n", restaurante1, calificacion1)
	fmt.Printf("El restaurante %s tiene  una calificación de %v \n", restaurante2, calificacion2)
	fmt.Printf("El restaurante %s tiene  una calificación de %v \n", restaurante3, calificacion3)
}
