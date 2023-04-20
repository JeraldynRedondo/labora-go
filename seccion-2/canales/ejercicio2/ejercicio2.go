/*
**Ejercicio 2**: Multiplicación de matrices utilizando gorutinas y canales
Dadas dos matrices A y B, crea un programa en Go que realice la multiplicación de las matrices utilizando gorutinas
y canales para dividir el trabajo entre varias gorutinas. Las matrices se pueden representar como matrices
bidimensionales (slices) en Go.
*/
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func MultiplicarMatrices(A, B [][]int, canal chan [][]int, wg *sync.WaitGroup) {
	defer wg.Done()

	if len(A[0]) != len(B) {
		fmt.Println("¡Inválido! Las matrices no pueden ser multiplicadas, verifique los datos")
		close(canal)
	} else {
		filas := len(A)
		columnas := len(B[0])

		C := make([][]int, filas)
		for i := range C {
			C[i] = make([]int, columnas)
		}

		for i := 0; i < filas; i++ {
			for j := 0; j < columnas; j++ {
				for k := 0; k < len(B); k++ {
					C[i][j] += A[i][k] * B[k][j]
				}
			}
		}
		canal <- C
	}

}
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // Configura el número máximo de núcleos que se pueden utilizar para ejecutar goroutines.
	var wg sync.WaitGroup

	//A, B := [][]int{{1, 2, 3, 4, 5}, {5, 6, 7, 8, 9}}, [][]int{{5, 4, 3, 2, 1}, {9, 8, 7, 6, 5}}
	A, B := [][]int{{2, 3}, {4, 1}}, [][]int{{5, 6}, {7, 8}}
	canal := make(chan [][]int)
	wg.Add(1)
	go MultiplicarMatrices(A, B, canal, &wg)
	C := <-canal
	wg.Wait()
	fmt.Println(C)

}
