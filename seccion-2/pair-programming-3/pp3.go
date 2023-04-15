package main

import (
	"fmt"
	"math"
)

// ///////INTERFACE/////////////
type Shapes interface {
	Area() float64
}

// /////////SHAPES/////////////
type Cuadrado struct {
	Lado float64
}
type Rectangulo struct {
	Base, Altura float64
}
type Triangulo struct {
	Base, Altura float64
}
type Rombo struct {
	Diagonal1, Diagonal2 float64
}
type Trapecio struct {
	Base1, Base2, Altura float64
}
type Circulo struct {
	Radio float64
}
type Pentagono struct {
	Lado, Apotema float64
}

// /////////FUNCIONES DE IMPLEMENTACIÓN INTERFACE/////////////
func (s *Cuadrado) Area() float64 {
	A := s.Lado * s.Lado
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("La operación falló:", err)
		}
	}()
	if s.Lado <= 0 {
		panic("Los números ingresados no pueden ser menor o igual a cero ")
	}
	if A >= 4000 {
		panic("¡El área es MUY grande,intentalo de nuevo! ")
	}
	return A
}
func (s *Rectangulo) Area() float64 {
	A := s.Base * s.Altura
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("La operación falló:", err)
		}
	}()
	if s.Altura <= 0 || s.Base <= 0 {
		panic("Los números ingresados no pueden ser menor o igual a cero ")
	}
	if s.Base == s.Altura {
		panic("Para que sea un rectangulo la base y la altura debe de ser distinta ")
	}
	if A >= 4000 {
		panic("¡El área es MUY grande,intentalo de nuevo! ")
	}
	return A
}
func (s *Triangulo) Area() float64 {
	A := (s.Base * s.Altura) / 2
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("La operación falló:", err)
		}
	}()
	if s.Altura <= 0 && s.Base <= 0 {
		panic("Los números ingresados no pueden ser menor o igual a cero ")
	}
	if A >= 4000 {
		panic("¡El área es MUY grande,intentalo de nuevo! ")
	}
	return A
}
func (s *Rombo) Area() float64 {
	A := (s.Diagonal1 * s.Diagonal2) / 2
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("La operación falló:", err)
		}
	}()
	if s.Diagonal1 <= 0 && s.Diagonal2 <= 0 {
		panic("Los números ingresados no pueden ser menor o igual a cero ")
	}
	if s.Diagonal1 == s.Diagonal2 {
		panic("Una diagonal debe de ser mayor a la otra, no iguales ")
	}
	if A >= 4000 {
		panic("¡El área es MUY grande,intentalo de nuevo! ")
	}
	return A
}
func (s *Trapecio) Area() float64 {
	A := ((s.Base1 + s.Base2) / 2) * s.Altura
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("La operación falló:", err)
		}
	}()
	if s.Base1 <= 0 && s.Base2 <= 0 && s.Altura <= 0 {
		panic("Los números ingresados no pueden ser menor o igual a cero ")
	}
	if s.Base1 == s.Base2 {
		panic("Para que sea un trapecio las dos bases deben de ser distintas ")
	}
	if A >= 4000 {
		panic("¡El área es MUY grande,intentalo de nuevo! ")
	}
	return A
}
func (s *Circulo) Area() float64 {
	A := math.Pi * s.Radio * s.Radio
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("La operación falló:", err)
		}
	}()
	if s.Radio <= 0 {
		panic("Los números ingresados no pueden ser menor o igual a cero ")
	}
	if A >= 4000 {
		panic("¡El área es MUY grande,intentalo de nuevo! ")
	}
	return A
}
func (s *Pentagono) Area() float64 {
	A := ((5 * s.Lado) * s.Apotema) / 2
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("La operación falló:", err)
		}
	}()
	if s.Lado <= 0 && s.Apotema <= 0 {
		panic("Los números ingresados no pueden ser menor o igual a cero ")
	}
	if s.Lado < s.Apotema {
		panic("El lado no puede ser menor que el apotema")
	}
	if A >= 4000 {
		panic("¡El área es MUY grande,intentalo de nuevo! ")
	}
	return A
}

///////////FUNCION DEL ÁREA//////////////

// /////////FUNCIONES DE USUARIO/////////////
func ImprimirAreas(figures map[string]Shapes) {
	for name, value := range figures {
		AREA := value.Area()
		if AREA != 0 && AREA < 4000 {
			fmt.Printf("El área del %s es: %v\n", name, AREA)
		}
	}
}
func main() {
	figures := map[string]Shapes{
		"Cuadrado":   &Cuadrado{Lado: 64},
		"Rectangulo": &Rectangulo{Base: 2, Altura: 4},
		"Triangulo":  &Triangulo{Base: 3, Altura: 6},
		"Rombo":      &Rombo{Diagonal1: 4, Diagonal2: 6},
		"Trapecio":   &Trapecio{Base1: 2, Base2: 2, Altura: 3},
		"Circulo":    &Circulo{Radio: 2},
		"Pentagono":  &Pentagono{Lado: 60, Apotema: 45},
	}

	ImprimirAreas(figures)
}
