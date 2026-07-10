package main

import "fmt"

func main() {
	// 1. Crear slice con literal
	nombres := []string{"Ana", "Luis", "Carlos"}
	fmt.Println("Nombres:", nombres)

	// 2. Slice a partir de un array
	numeros := [5]int{10, 20, 30, 40, 50}
	slice := numeros[1:4]
	fmt.Println("Array completo:", numeros)
	fmt.Println("Slice [1:4]:", slice)

	// 3. Slice sin índice inicial ni final
	todo := numeros[:]
	fmt.Println("Slice numeros[:]:", todo)

	// 4. Slice con make (tipo, longitud, capacidad)
	valores := make([]int, 3, 5)
	fmt.Println("Slice con make:", valores)
	fmt.Println("Longitud:", len(valores), "Capacidad:", cap(valores))

	// 5. Slice vacío vs nil
	var vacio []int
	fmt.Println("Slice nil:", vacio, "len:", len(vacio), "nil?", vacio == nil)

	vacio2 := []int{}
	fmt.Println("Slice vacío:", vacio2, "len:", len(vacio2), "nil?", vacio2 == nil)

	// 6. Rebanar un slice
	superior := []int{1, 2, 3, 4, 5, 6}
	mitad := superior[2:5]
	fmt.Println("Slice original:", superior)
	fmt.Println("Rebanada [2:5]:", mitad)

	fmt.Scanln()
}
