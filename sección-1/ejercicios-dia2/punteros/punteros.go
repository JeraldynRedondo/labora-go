/*
Escribir un programa en Go que declare dos variables enteras a y b con valores iniciales de 10 y 20, respectivamente.
A continuación, declarar un puntero ptrA que apunte a la dirección de memoria de a. Utilizar el puntero ptrA para
intercambiar los valores de a y b sin utilizar una variable auxiliar. Finalmente, imprimir los valores de a y b
por pantalla para comprobar que se han intercambiado correctamente.
*/
package main

import (
	"fmt"
)

func main() {
	var a, b int = 10, 20
	var ptrA *int = &a

	fmt.Printf("Valores iniciales: a = %d , b = %d  \n", a, b)

	a = b + *ptrA
	b = a - b
	a = *ptrA - b

	fmt.Printf("Valores intercambiados: a = %d , b = %d  \n", a, b)
}
