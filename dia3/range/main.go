package main

import "fmt"

func main() {
	// 1. range sobre slice (índice y valor)
	frutas := []string{"manzana", "pera", "uva", "naranja"}
	fmt.Println("=== Frutas (índice + valor) ===")
	for i, fruta := range frutas {
		fmt.Printf("%d: %s\n", i, fruta)
	}

	// 2. range ignorando el índice con _
	fmt.Println("\n=== Solo valores ===")
	for _, fruta := range frutas {
		fmt.Println(fruta)
	}

	// 3. range ignorando el valor
	fmt.Println("\n=== Solo índices ===")
	for i := range frutas {
		fmt.Println("índice:", i)
	}

	// 4. range sobre array
	numeros := [5]int{10, 20, 30, 40, 50}
	fmt.Println("\n=== Array ===")
	for i, v := range numeros {
		fmt.Printf("numeros[%d] = %d\n", i, v)
	}

	// 5. range modifica el slice original (no copia)
	fmt.Println("\n=== Modificar con range (por índice) ===")
	cuadrados := []int{1, 2, 3, 4, 5}
	for i := range cuadrados {
		cuadrados[i] = cuadrados[i] * cuadrados[i]
	}
	fmt.Println("Cuadrados:", cuadrados)

	// 6. range con string (recorre runas, no bytes)
	fmt.Println("\n=== range sobre string ===")
	palabra := "Go ñoño"
	for i, r := range palabra {
		fmt.Printf("byte %d -> %c (%U)\n", i, r, r)
	}

	// 7. range sobre map (solo para mostrar, se verá en día 4)
	edades := map[string]int{"Ana": 25, "Luis": 30}
	fmt.Println("\n=== range sobre map ===")
	for nombre, edad := range edades {
		fmt.Printf("%s tiene %d años\n", nombre, edad)
	}

	fmt.Scanln()
}
