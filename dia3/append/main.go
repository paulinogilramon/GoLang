package main

import "fmt"

func main() {
	// 1. append a un slice
	var numeros []int
	numeros = append(numeros, 10)
	numeros = append(numeros, 20, 30, 40)
	fmt.Println("Slice después de append:", numeros)

	// 2. append con slice existente
	pares := []int{2, 4, 6}
	impares := []int{1, 3, 5}
	todos := append(pares, impares...)
	fmt.Println("Pares:", pares)
	fmt.Println("Impares:", impares)
	fmt.Println("Combinados:", todos)

	// 3. append y capacidad
	fmt.Println("\n=== Crecimiento de capacidad ===")
	var nums []int
	for i := 1; i <= 10; i++ {
		nums = append(nums, i)
		fmt.Printf("len: %d, cap: %d, valor: %v\n", len(nums), cap(nums), nums)
	}

	// 4. append con make preasignado
	datos := make([]string, 0, 3)
	fmt.Println("\nAntes de append:", datos, "len:", len(datos), "cap:", cap(datos))
	datos = append(datos, "a", "b", "c", "d")
	fmt.Println("Después de append:", datos, "len:", len(datos), "cap:", cap(datos))

	// 5. append y rebanado
	original := []int{1, 2, 3, 4, 5}
	nuevo := append(original[:2], original[3:]...)
	fmt.Println("\nEliminar índice 2:", nuevo)

	fmt.Scanln()
}
