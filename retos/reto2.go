/*
Pasar un periodo expresado en segundos a un periodo expresado en días, horas, minutos y segundos.
1030 segundos son 17 minutos con 10 segundos
12045 segundos son 3 horas, 20 minutos con 45 segundos
176520 segundos son 2 días, 1 hora, 2 minutos con 0 segundos.
*/
package main

import (
	"fmt"
)

func CalcularTiempo(segundos int) {
	num := segundos
	dias := segundos / 86400
	segundos %= 86400
	horas := segundos / 3600
	segundos %= 3600
	minutos := segundos / 60
	segundos %= 60

	fmt.Printf("%d Segundos son %d dias, %d horas, %d minutos y %d segundos \n", num, dias, horas, minutos, segundos)
}

func main() {
	//var num int
	//fmt.Println("¡Bienvenido! \nPor favor,ingresa el número de segundos a calcular= ")
	//fmt.Scanln(&num)
	//CalcularTiempo(num)

	CalcularTiempo(80)
	CalcularTiempo(7200)
	CalcularTiempo(10820)
	CalcularTiempo(12000)
	CalcularTiempo(2000)
	CalcularTiempo(86400)
	CalcularTiempo(87450)

}
