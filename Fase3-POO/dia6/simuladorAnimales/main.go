package main

import "fmt"

type Animal interface {
	HacerSonido() string
	Moverse() string
	Comer() string
}

type Perro struct{ nombre string }

func (p Perro) HacerSonido() string { return "Guau guau!" }
func (p Perro) Moverse() string     { return "Corre en cuatro patas" }
func (p Perro) Comer() string       { return "Come croquetas" }

type Pajaro struct{ nombre string }

func (p Pajaro) HacerSonido() string { return "Pío pío!" }
func (p Pajaro) Moverse() string     { return "Vuela por el aire" }
func (p Pajaro) Comer() string       { return "Come semillas" }

type Pez struct{ nombre string }

func (p Pez) HacerSonido() string { return "Glub glub" }
func (p Pez) Moverse() string     { return "Nada en el agua" }
func (p Pez) Comer() string       { return "Come alimento para peces" }

type Embedding struct{}

func (p Embedding) HacerSonido() string { return "..." }

// Embedding de struct (composición)
type AnimalConNombre struct {
	Animal
	nombre string
}

func (a AnimalConNombre) Nombre() string {
	return a.nombre
}

func main() {
	animales := []Animal{
		Perro{"Rex"},
		Pajaro{"Tweety"},
		Pez{"Nemo"},
	}

	// Simular animales con sus comportamientos
	for _, a := range animales {
		fmt.Printf("\n%s\n  Sonido: %s\n  Movimiento: %s\n  Comida: %s\n",
			obtenerNombre(a), a.HacerSonido(), a.Moverse(), a.Comer())
	}

	// Embedding de interfaz
	aguila := AnimalConNombre{
		Animal: Pajaro{"Águila"},
		nombre: "Águila Real",
	}
	fmt.Printf("\n%s: %s\n", aguila.Nombre(), aguila.HacerSonido())
}

func obtenerNombre(a Animal) string {
	switch v := a.(type) {
	case Perro:
		return v.nombre
	case Pajaro:
		return v.nombre
	case Pez:
		return v.nombre
	default:
		return "Animal"
	}
}
