/*
Crea un programa en Go que tome un número como entrada y devuelva el nombre del día de la semana correspondiente.
Por ejemplo, si se ingresa "1", debe devolver "Lunes".
*/
package main

import (
	"fmt"
)

func main() {
	var eleccion int
	fmt.Println("¡Bienvenido! \nPor favor,ingresa un número del 1 al 7")
	fmt.Scanf("%d\n", &eleccion)
	switch eleccion {
	case 1:
		fmt.Printf("El día de la semana que coincide con el número %d es Lunes", eleccion)
	case 2:
		fmt.Printf("El día de la semana que coincide con el número %d es Martes", eleccion)
	case 3:
		fmt.Printf("El día de la semana que coincide con el número %d es Miércoles", eleccion)
	case 4:
		fmt.Printf("El día de la semana que coincide con el número %d es Jueves", eleccion)
	case 5:
		fmt.Printf("El día de la semana que coincide con el número %d es Viernes", eleccion)
	case 6:
		fmt.Printf("El día de la semana que coincide con el número %d es Sábado", eleccion)
	case 7:
		fmt.Printf("El día de la semana que coincide con el número %d es Domingo", eleccion)
	default:
		fmt.Printf("El número %d, no corresponde a ningún día de la semana.\nPor favor intente de nuevo", eleccion)
	}
}
