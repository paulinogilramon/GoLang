package main

import "fmt"

// Struct básico
type Persona struct {
	Nombre string
	Edad   int
}

// Receiver por valor
func (p Persona) Saludar() {
	fmt.Printf("Hola, soy %s y tengo %d años\n", p.Nombre, p.Edad)
}

// Receiver por puntero (modifica el struct)
func (p *Persona) CumplirAnios() {
	p.Edad++
}

// Struct con campos anidados
type Direccion struct {
	Calle string
	Numero int
	Ciudad string
}

type Usuario struct {
	Persona
	Email    string
	Direccion Direccion
}

// Método con receiver por valor
func (u Usuario) MostrarInfo() {
	fmt.Printf("Usuario: %s (%s)\n", u.Nombre, u.Email)
	fmt.Printf("Dirección: %s %d, %s\n", u.Direccion.Calle, u.Direccion.Numero, u.Direccion.Ciudad)
}

func main() {
	// Crear instancias
	p1 := Persona{"Alice", 30}
	p1.Saludar()
	p1.CumplirAnios()
	p1.Saludar()

	// Constructor con puntero
	p2 := &Persona{"Bob", 25}
	p2.Saludar()
	p2.CumplirAnios()
	p2.Saludar()

	// Struct con embedding y campos anidados
	u := Usuario{
		Persona: Persona{"Charlie", 28},
		Email:   "charlie@example.com",
		Direccion: Direccion{
			Calle:  "Av. Siempre Viva",
			Numero: 742,
			Ciudad: "Springfield",
		},
	}
	u.MostrarInfo()

	// Zero value de struct
	var p3 Persona
	fmt.Printf("Zero value: %+v\n", p3)

	// Struct anónimo
	anonimo := struct {
		Nombre string
		Edad   int
	}{"Anónimo", 99}
	fmt.Println("Struct anónimo:", anonimo)
}
