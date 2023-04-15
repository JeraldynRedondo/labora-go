/*
1 - En nuestra carpeta de la sección 1, crear una carpeta llamada **arrays**, dentro crear un archivo `**array.go`**
en donde se haga:
- Crear una matriz con el número 0 a 10 y mostrarlos en pantalla.
- Crear una matriz de cadenas con nombres y mostrarlos en pantalla.
*/
package main

import (
	"fmt"
)

func main() {

	var numeros = [11]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var nombres = [5]string{"Jeraldyn", "Elena", "Sheryl", "Emily", "Analia"}

	fmt.Println("La matriz de números es: ", numeros)
	fmt.Println("La matriz de nombres es: ", nombres)

}
