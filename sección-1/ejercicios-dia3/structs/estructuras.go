/*
3. Define una estructura llamada **`Persona`** con los siguientes campos: **`nombre`**, **`edad`**,
**`ciudad`** y **`telefono`**.
4. Crea dos variables de tipo **`Persona`** con valores iniciales distintos.
5. Imprime los valores de ambas variables por pantalla.
6. Define una función llamada **`cambiarCiudad`** que reciba una persona y una ciudad como parámetros y actualice
el campo **`ciudad`** de la persona solo si la ciudad es distinta a la actual. Si la ciudad es la misma, no hace nada.
7. Llama a la función **`cambiarCiudad`** con una de las personas creadas anteriormente y una ciudad distinta a la actual.
8. Imprime los valores de ambas variables por pantalla para comprobar que se ha actualizado el campo **`ciudad`** de
la persona correspondiente si la ciudad es distinta, o que no se ha actualizado si la ciudad es la misma.
*/
package main

import (
	"fmt"
)

type Persona struct {
	nombre   string
	edad     int
	ciudad   string
	telefono string
}

func cambiarCiudad(persona *Persona, ciudad string) {
	if persona.ciudad != ciudad {
		persona.ciudad = ciudad
	}

}

func main() {

	persona1 := Persona{nombre: "Juan", edad: 30, ciudad: "Madrid", telefono: "555-1234"}
	persona2 := Persona{nombre: "Ana", edad: 25, ciudad: "Barcelona", telefono: "555-5678"}

	fmt.Println("Persona 1: ", persona1)
	fmt.Println("Persona 2: ", persona2)

	cambiarCiudad(&persona1, "Valencia")
	fmt.Println("Persona 1 con ciudad actualizada: ", persona1)
	fmt.Println("Persona 2: ", persona2)

}
