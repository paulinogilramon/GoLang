package main

import "fmt"

func main() {
	// 1. Declarar array con tamaño fijo
	var numeros [5]int
	fmt.Println("Array sin inicializar:", numeros)

	// 2. Asignar valores por índice
	numeros[0] = 10
	numeros[1] = 20
	numeros[4] = 50
	fmt.Println("Array con valores:", numeros)

	// 3. Array literal
	frutas := [3]string{"manzana", "pera", "uva"}
	fmt.Println("Frutas:", frutas)

	// 4. Inferir tamaño con [...]
	pares := [...]int{2, 4, 6, 8, 10}
	fmt.Println("Pares:", pares)

	// 5. Acceder por índice
	fmt.Println("Primera fruta:", frutas[0])
	fmt.Println("Último par:", pares[len(pares)-1])

	// 6. Modificar un elemento
	frutas[1] = "naranja"
	fmt.Println("Frutas modificado:", frutas)

	// 7. Longitud con len
	fmt.Println("Tamaño de pares:", len(pares))

	fmt.Scanln()
}
