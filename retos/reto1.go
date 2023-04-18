/*
Expresar a un número cualquiera como la suma de una serie de números siguiendo las restricciones impuestas a continuación.

Sea x el número.
x = s1 + s2 + s3 + s4 + s5

Donde
0 < s1 < 50
0 < s2 < 50
0 < s3 < 600
0 < s4 < 800
0 < s5 < (Infinito)

De esta forma, si X vale 120, entonces
s1 = 50, s2 = 50, s3 = 20, s4 = 0 y s5 = 0

Si X vale 860, entonces
s1 = 50, s2 = 50, s3 = 600, s4 = 160 y s5 = 0
*/
package main

import (
	"fmt"
)

func CalcularNumeros(num int) {
	var s1, s2, s3, s4, s5 int

	if num <= 50 {
		s1 = 50

		s2, s3, s4, s5 = 0, 0, 0, 0
	} else if num <= 100 {
		s1 = 50
		s2 = num - 50

	} else if num <= 700 {
		s1, s2 = 50, 50
		s3 = num - s1 - s2
		s4, s5 = 0, 0
	} else if num <= 1500 {
		s1, s2, s3 = 50, 50, 600
		s4 = num - s1 - s2 - s3
		s5 = 0
	} else {
		s1 = 50
		s2 = 50
		s3 = 600
		s4 = 800
		s5 = num - 1500
	}

	fmt.Printf("Si X vale %d entonces \n s1 = %d, s2 = %d, s3 = %d, s4 = %d y s5 = %d \n", num, s1, s2, s3, s4, s5)
}

func main() {
	//var num int
	//fmt.Println("¡Bienvenido! \nPor favor,ingresa el número= ")
	//fmt.Scanf("%d\n", &num)

	CalcularNumeros(80)
	CalcularNumeros(120)
	CalcularNumeros(400)
	CalcularNumeros(860)
	CalcularNumeros(2000)

}
