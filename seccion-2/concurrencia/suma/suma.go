package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func SumarNumerosSecuencial(numeros []int) int {
	var suma int
	for _, value := range numeros {
		suma += value
	}
	return suma
}

func SumarNumerosConcurrente(numeros []int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	var suma int
	for _, value := range numeros {
		suma += value
	}

	c <- suma

}

func main() {
	numeros := make([]int, 1000000)
	//Llenar vector

	for i := 0; i < len(numeros); i++ {
		numeros[i] = i + 1
	}

	start := time.Now()
	//Secuencial
	Suma := SumarNumerosSecuencial(numeros)
	fmt.Printf("La suma de los primeros %d números es igual a= %d \n", len(numeros), Suma)
	durationS := time.Since(start)
	fmt.Println("La duración de la forma secuencial es: ", durationS)

	//Concurrente
	runtime.GOMAXPROCS(runtime.NumCPU()) // Configura el número máximo de núcleos que se pueden utilizar para ejecutar goroutines.

	var wg sync.WaitGroup
	c := make(chan int)
	rango := len(numeros) / 2
	wg.Add(2)

	start2 := time.Now()
	go SumarNumerosConcurrente((numeros[:rango]), &wg, c)
	go SumarNumerosConcurrente((numeros[rango:]), &wg, c)
	suma1, suma2 := <-c, <-c
	SumaTotal := suma1 + suma2
	fmt.Printf("La suma de los primeros %d números es igual a= %d \n", len(numeros), SumaTotal)
	wg.Wait()
	durationC := time.Since(start2)
	fmt.Println("La duración de la forma Concurrente es: ", durationC)
}
