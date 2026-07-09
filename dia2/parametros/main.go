package main

import "fmt"

// 1. Paso por valor — no modifica el original
func incrementarValor(n int) {
	n++
}

// 2. Paso por referencia (puntero) — modifica el original
func incrementarPuntero(n *int) {
	*n++
}

// 3. Parámetros variádicos
func sumar(numeros ...int) int {
	total := 0
	for _, n := range numeros {
		total += n
	}
	return total
}

// 4. Slice como parámetro — se modifica el contenido interno
func duplicar(nums []int) {
	for i := range nums {
		nums[i] *= 2
	}
}

// 5. Map como parámetro — se agrega/ modifica directamente
func agregarEdad(m map[string]int, nombre string, edad int) {
	m[nombre] = edad
}

// 6. Función como parámetro (callback)
func aplicarOperacion(a, b int, op func(int, int) int) int {
	return op(a, b)
}

func main() {
	// 1. Paso por valor
	x := 10
	incrementarValor(x)
	fmt.Println("1. Paso por valor — x sigue siendo:", x)

	// 2. Paso por puntero
	y := 10
	incrementarPuntero(&y)
	fmt.Println("2. Paso por puntero — y ahora es:", y)

	// 3. Variádicos
	fmt.Println("3. Suma variádica:", sumar(1, 2, 3, 4, 5))
	fmt.Println("   Suma sin argumentos:", sumar())

	// 4. Slice como parámetro
	numeros := []int{1, 2, 3, 4}
	duplicar(numeros)
	fmt.Println("4. Slice después de duplicar:", numeros)

	// 5. Map como parámetro
	edades := make(map[string]int)
	agregarEdad(edades, "Ana", 30)
	agregarEdad(edades, "Luis", 25)
	fmt.Println("5. Map después de agregar edades:", edades)

	// 6. Función como parámetro
	resultado := aplicarOperacion(10, 5, func(a, b int) int {
		return a + b
	})
	fmt.Println("6. Resultado del callback (10+5):", resultado)

	fmt.Println("\nPresiona Enter para salir...")
	fmt.Scanln()
}
