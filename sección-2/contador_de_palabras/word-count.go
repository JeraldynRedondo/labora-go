package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	v := strings.Fields(s)
	conteo := 0
	palabras := make(map[string]int)
	for _, value := range v {
		_, ok := palabras[value]
		if !ok {
			conteo = 0
			for j := 0; j < len(v); j++ {
				if value == v[j] {
					conteo++
				}
			}
			palabras[value] = conteo
		}
	}
	return palabras
}

func main() {
	resultado := WordCount("I ate a donut. Then I ate another donut.")
	fmt.Println(resultado)
}
