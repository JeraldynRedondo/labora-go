/*
Hacer un programa que lance dos go rutinas
En una haga entrada de teclado
Y otra que cree un arreglo de muchos numeros
y que sume esos numeros
al final se imprime el string y la suma
*/
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func Entrada(c chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	var Entrada string
	fmt.Println("¡Bienvenido! \nIngrese lo que desee: ")
	fmt.Scanln(&Entrada)

	c <- Entrada
}

func SumarNumeros(c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	suma := 0
	for i := 0; i < 100000000; i++ {
		suma = i + 1
	}

	c <- suma
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // Configura el número máximo de núcleos que se pueden utilizar para ejecutar goroutines.

	var wg sync.WaitGroup
	cI := make(chan int)
	cS := make(chan string)

	wg.Add(2)

	go Entrada(cS, &wg)
	go SumarNumeros(cI, &wg)

	salida, numeros := <-cS, <-cI

	fmt.Printf("La entrada recibida fue: %s \n", salida)
	fmt.Printf("La suma es igual a= %d \n", numeros)

	wg.Wait()
	close(cI)
	close(cS)

}
