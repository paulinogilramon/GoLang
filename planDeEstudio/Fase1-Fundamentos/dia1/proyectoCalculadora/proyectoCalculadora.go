package main

import "fmt"

func main() {
	const (
		Salir = iota
		Sumar
		Restar
		Multiplicar
		Dividir
	)

	var opcion int
	var num1, num2, resultado float64

	for {
		fmt.Println("\n=== CALCULADORA ===")
		fmt.Println("1. Sumar")
		fmt.Println("2. Restar")
		fmt.Println("3. Multiplicar")
		fmt.Println("4. Dividir")
		fmt.Println("0. Salir")
		fmt.Print("Elige una opción: ")
		fmt.Scan(&opcion)

		if opcion == Salir {
			fmt.Println("¡Hasta luego!")
			fmt.Scanln()
			break
		}

		if opcion < Sumar || opcion > Dividir {
			fmt.Println("Opción no válida")
			continue
		}

		fmt.Print("Ingresa el primer número: ")
		fmt.Scan(&num1)
		fmt.Print("Ingresa el segundo número: ")
		fmt.Scan(&num2)

		switch opcion {
		case Sumar:
			resultado = num1 + num2
			fmt.Printf("%.2f + %.2f = %.2f\n", num1, num2, resultado)
		case Restar:
			resultado = num1 - num2
			fmt.Printf("%.2f - %.2f = %.2f\n", num1, num2, resultado)
		case Multiplicar:
			resultado = num1 * num2
			fmt.Printf("%.2f * %.2f = %.2f\n", num1, num2, resultado)
		case Dividir:
			if num2 == 0 {
				fmt.Println("Error: no se puede dividir entre cero")
				continue
			}
			resultado = num1 / num2
			fmt.Printf("%.2f / %.2f = %.2f\n", num1, num2, resultado)
		}
	}
}
