package main

import "fmt"

func main() {
	// 1. if simple
	edad := 18
	if edad >= 18 {
		fmt.Println("Eres mayor de edad")
	}

	// 2. if-else
	temperatura := 30
	if temperatura > 25 {
		fmt.Println("Hace calor")
	} else {
		fmt.Println("No hace tanto calor")
	}

	// 3. if-else if-else
	nota := 85
	if nota >= 90 {
		fmt.Println("Sobresaliente")
	} else if nota >= 70 {
		fmt.Println("Notable")
	} else if nota >= 50 {
		fmt.Println("Aprobado")
	} else {
		fmt.Println("Reprobado")
	}

	// 4. if con short statement (declaración + condición)
	if x := 10; x > 5 {
		fmt.Println("x es mayor que 5")
	}

	// 5. Operadores lógicos (&&, ||, !)
	a, b := 7, 12
	if a > 5 && b < 20 {
		fmt.Println("Ambas condiciones se cumplen")
	}
	if a > 10 || b < 20 {
		fmt.Println("Al menos una se cumple")
	}

	fmt.Scanln()
}
