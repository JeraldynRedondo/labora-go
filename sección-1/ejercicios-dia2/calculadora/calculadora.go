/*
Escribir un programa que pida al usuario que ingrese dos números enteros. Luego, se debe crear una función llamada
calcular que tome dos punteros a enteros como argumentos y calcule la suma, resta, multiplicación y división de los

	números apuntados por los punteros. Finalmente, se deben imprimir los resultados de las operaciones por pantalla.
*/
package main

import (
	"fmt"
)

func calcular(a, b *int) {
	var suma = *a + *b
	var resta = *a - *b
	var multi = (*a) * (*b)
	var div = (*a) / (*b)

	fmt.Printf("La suma de a=%d y b=%d = %d \n", *a, *b, suma)
	fmt.Printf("La resta de a=%d y b=%d = %d \n", *a, *b, resta)
	fmt.Printf("La multiplicación de a=%d y b=%d = %d \n", *a, *b, multi)
	fmt.Printf("La división de a=%d y b=%d = %d \n", *a, *b, div)
}

func main() {
	var a, b int
	fmt.Println("¡Bienvenido! \nPor favor,ingresa el primer número= ")
	fmt.Scanf("%d\n", &a)
	fmt.Println("Por favor,ingresa el segundo número= ")
	fmt.Scanf("%d\n", &b)

	fmt.Printf("Sus números son a=%d y b=%d \n", a, b)
	calcular(&a, &b)

}
