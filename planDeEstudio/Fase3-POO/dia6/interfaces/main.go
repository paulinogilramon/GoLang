package main

import "fmt"

// Definición de interfaces
type Animal interface {
	Sonido() string
	Nombre() string
}

// Interface con múltiples métodos
type Trabajador interface {
	Trabajar() string
	Descansar() string
}

// Interface embedida
type Mascota interface {
	Animal
	Jugar() string
}

// --- Implementaciones de Animal ---

type Perro struct {
	nombre string
}

func (p Perro) Sonido() string { return "Guau" }
func (p Perro) Nombre() string { return p.nombre }

type Gato struct {
	nombre string
}

func (g Gato) Sonido() string { return "Miau" }
func (g Gato) Nombre() string { return g.nombre }

type Vaca struct {
	nombre string
}

func (v Vaca) Sonido() string { return "Muu" }
func (v Vaca) Nombre() string { return v.nombre }

// --- Interface vacía (any) ---

func Describir(i interface{}) {
	fmt.Printf("Valor: %v, Tipo: %T\n", i, i)
}

// Type assertion
func IdentificarAnimal(a Animal) {
	switch a.(type) {
	case Perro:
		fmt.Println("Es un perro!")
	case Gato:
		fmt.Println("Es un gato!")
	case Vaca:
		fmt.Println("Es una vaca!")
	default:
		fmt.Println("Es otro animal")
	}

	// Type assertion con ok
	if perro, ok := a.(Perro); ok {
		fmt.Printf("Es un perro llamado %s\n", perro.nombre)
	}
}

func main() {
	// Usar la interfaz Animal
	animales := []Animal{
		Perro{"Rex"},
		Gato{"Misi"},
		Vaca{"Lola"},
	}

	for _, a := range animales {
		fmt.Printf("%s hace: %s\n", a.Nombre(), a.Sonido())
	}

	// Type assertion e identificación
	fmt.Println()
	IdentificarAnimal(Perro{"Max"})
	IdentificarAnimal(Gato{"Garfield"})

	// Interface vacía
	fmt.Println()
	Describir(42)
	Describir("hola")
	Describir(3.14)
	Describir(Perro{"Fido"})
}
