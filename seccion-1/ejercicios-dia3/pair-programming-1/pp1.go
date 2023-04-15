/*
 */
package main

import (
	"fmt"
)

func main() {

	persona1 := Persona{nombre: "Juan", edad: 30, ciudad: "Madrid", telefono: 5551234}
	persona2 := Persona{nombre: "Ana", edad: 25, ciudad: "Barcelona", telefono: 5555678}
	persona3 := Persona{nombre: "Jose", edad: 20, ciudad: "Sevilla", telefono: 5559876}
	persona4 := Persona{nombre: "Maria", edad: 25, ciudad: "Barcelona", telefono: 5556789}
	persona5 := Persona{nombre: "Sofia", edad: 24, ciudad: "Toledo", telefono: 5551478}

	incrementarEdad(&persona1)
	fmt.Println("La nueva edad de la persona es: ", persona1.edad)

	personas := [5]Persona{persona1, persona2, persona3, persona4, persona5}

	b := buscarEdad(personas, 25)

	if b != nil {
		fmt.Printf("¡Encontrada! La persona %s tiene la edad de %d años \n", b.nombre, b.edad)
	} else {
		fmt.Printf("No es posible encontrar una persona con esa edad,intentelo de nuevo")
	}

	personaNueva := crearPersona("Diego", 23, "Madrid", 5551454)

	fmt.Println("¡Creación exitosa! la nueva persona tiene estas características", personaNueva)

	actualizarTelefono(&personaNueva, 55551474)

	fmt.Printf("El numero de teléfono de %s ha sido actualizado de manera exitosa y es: %d \n", personaNueva.nombre, personaNueva.telefono)

}

type Persona struct {
	nombre   string
	edad     int
	ciudad   string
	telefono int64
}

func incrementarEdad(persona *Persona) {
	persona.edad++
}

func buscarEdad(personas [5]Persona, edad int) *Persona {
	for i := 0; i < len(personas); i++ {
		if personas[i].edad == edad {
			return &personas[i]
		}
	}
	return nil
}

func crearPersona(nombre string, edad int, ciudad string, telefono int64) Persona {
	persona := Persona{nombre: nombre, edad: edad, ciudad: ciudad, telefono: telefono}
	return persona
}

func actualizarTelefono(persona *Persona, telefono int64) {
	persona.telefono = telefono
}
