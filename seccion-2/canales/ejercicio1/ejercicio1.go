/*
**Ejercicio 1**: Suma de números en un rango utilizando gorutinas y canales
Escribe un programa en Go que sume todos los números en un rango dado (por ejemplo, de 1 a 100) utilizando
gorutinas y canales para dividir el trabajo entre varias gorutinas.
*/
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func SumarNumeros(Inicio, Final int, c chan int, wg *sync.WaitGroup) {
	suma := 0
	for i := Inicio; i <= Final; i++ {
		suma += i
	}
	c <- suma
	wg.Done()
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // Configura el número máximo de núcleos que se pueden utilizar para ejecutar goroutines.
	c := make(chan int)
	var (
		wg            sync.WaitGroup
		Inicio, Final int
	)

	// fmt.Println("¡Bienvenido! \nPor favor,ingresa el número de inicio del rango")
	// fmt.Scanln(&Inicio)

	// fmt.Println("Por favor,ingresa el número Final del rango")
	// fmt.Scanln(&Final)

	Inicio, Final = 1, 100
	NumGr := 10
	RangoGr := (Final - Inicio + 1) / NumGr

	for i := 0; i < NumGr; i++ {
		wg.Add(1)
		go SumarNumeros(Inicio+i*RangoGr, Inicio+(i+1)*RangoGr-1, c, &wg)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	SumaTotal := 0
	for suma := range c {
		SumaTotal += suma
	}

	fmt.Printf("La suma de los números del %d al %d es igual a= %d \n", Inicio, Final, SumaTotal)

}
