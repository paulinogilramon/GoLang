package main

import "fmt"

func main() {
	// var con tipo explícito
	var edad int = 25

	// var sin valor inicial (zero value: 0)
	var numero int

	// Inferencia de tipo con :=
	nombre := "Carlos"
	altura := 1.75

	// Múltiples variables
	var x, y int = 10, 20
	a, b := "hola", 42

	fmt.Println("=== Variables ===")
	fmt.Println("Nombre:", nombre)
	fmt.Println("Edad:", edad)
	fmt.Println("Altura:", altura)
	fmt.Println("Numero (zero value):", numero)
	fmt.Println("X:", x, "Y:", y)
	fmt.Println("A:", a, "B:", b)

	// Zero values de otros tipos
	var texto string
	var flotante float64
	var booleano bool
	fmt.Println("\n=== Zero Values ===")
	fmt.Println("Booleano:", booleano)
	fmt.Printf("String:%q\n", texto)
	fmt.Println("Float64:", flotante)

	fmt.Scanln()
}
