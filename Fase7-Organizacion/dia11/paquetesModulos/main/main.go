package main

import (
	"fmt"

	"github.com/rapaulinog/golang-estudio/Fase7-Organizacion/dia11/paquetesModulos/matematica"
)

func main() {
	// Usar funciones del paquete matematica
	a, b := 10.0, 3.0

	fmt.Printf("%.1f + %.1f = %.1f\n", a, b, matematica.Suma(a, b))
	fmt.Printf("%.1f - %.1f = %.1f\n", a, b, matematica.Resta(a, b))
	fmt.Printf("%.1f * %.1f = %.1f\n", a, b, matematica.Multiplica(a, b))

	resultado, err := matematica.Divide(a, b)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("%.1f / %.1f = %.1f\n", a, b, resultado)
	}

	// Prueba de division por cero
	resultado, err = matematica.Divide(5, 0)
	if err != nil {
		fmt.Println("Error controlado:", err)
	}

	fmt.Printf("2^8 = %.0f\n", matematica.Potencia(2, 8))
	fmt.Printf("Factorial de 5 = %d\n", matematica.Factorial(5))
	fmt.Printf("Fibonacci(10) = %d\n", matematica.Fibonacci(10))
}
