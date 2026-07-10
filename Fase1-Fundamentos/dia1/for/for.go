package main

import "fmt"

func main() {
	// 1. for clásico (inicialización; condición; incremento)
	fmt.Println("=== for clásico ===")
	for i := 0; i < 5; i++ {
		fmt.Println("i:", i)
	}

	// 2. for como while (solo condición)
	fmt.Println("\n=== for como while ===")
	suma := 0
	for suma < 5 {
		suma++
		fmt.Println("suma:", suma)
	}

	// 3. for infinito (se rompe con break)
	fmt.Println("\n=== for infinito con break ===")
	contador := 0
	for {
		if contador >= 3 {
			break
		}
		fmt.Println("contador:", contador)
		contador++
	}

	// 4. for range sobre un slice
	fmt.Println("\n=== for range (slice) ===")
	frutas := []string{"manzana", "pera", "uva"}
	for i, fruta := range frutas {
		fmt.Printf("índice %d -> %s\n", i, fruta)
	}

	// 5. for range ignorando el índice con _
	fmt.Println("\n=== for range (solo valor) ===")
	for _, fruta := range frutas {
		fmt.Println(fruta)
	}

	// 6. continue
	fmt.Println("\n=== continue (saltar pares) ===")
	for n := 1; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println("impar:", n)
	}

	fmt.Scanln()
}
