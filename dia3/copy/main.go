package main

import "fmt"

func main() {
	// 1. copy con slices del mismo tamaño
	origen := []int{1, 2, 3, 4, 5}
	destino := make([]int, 5)
	copied := copy(destino, origen)
	fmt.Println("Origen:", origen)
	fmt.Println("Destino:", destino)
	fmt.Println("Elementos copiados:", copied)

	// 2. copy con destino más pequeño (solo copia hasta len(destino))
	destinoChico := make([]int, 3)
	copy(destinoChico, origen)
	fmt.Println("\nDestino pequeño (len 3):", destinoChico)

	// 3. copy con origen más pequeño
	origenChico := []int{10, 20}
	destinoGrande := make([]int, 5)
	copy(destinoGrande, origenChico)
	fmt.Println("\nDestino grande, origen chico:", destinoGrande)

	// 4. copy entre slices que se solapan
	numeros := []int{1, 2, 3, 4, 5}
	copy(numeros[1:], numeros[0:])
	fmt.Println("\nSolapamiento [1:] desde [0:]:", numeros)

	// 5. copy con string a []byte
	var buf [8]byte
	copy(buf[:], "Hola Go")
	fmt.Printf("\nString a bytes: %v\n", buf)
	fmt.Printf("Como string: %q\n", string(buf[:]))

	// 6. copy no agranda el destino
	a := []int{1, 2}
	b := []int{100, 200, 300}
	n := copy(a, b)
	fmt.Println("\nOrigen más grande que destino:", a, "(copiados:", n, ")")

	fmt.Scanln()
}
