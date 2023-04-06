/*
2 - Crear un archivo llamado **peliculas.go** en donde tengas que crear una matriz llamada "películas" que contenga solo
nombres, de hasta 10 elementos. Imprime lo siguiente:
- La matriz en sí
- El segundo elemento de la matriz.
- La longitud de la matriz
*/
package main

import (
	"fmt"
)

func main() {

	películas := [10]string{"Titanic", "Ben-Hur", "El Señor de los Anillos: el retorno del Rey",
		"Lo que el viento se llevó", "West Side Story", "El paciente inglés", "Gigi", "El último emperador",
		"De aquí a la eternidad", "On the Waterfront"}

	fmt.Println("Las películas con más premios Oscar son: \n", películas)
	fmt.Printf("La película %s ganó 11 estatuilla del premio Oscar en la edición 32 de los premios en el año 1959 \n", películas[1])
	fmt.Printf("La matriz películas tiene una longitud de %d elementos \n", len(películas))
}
