package main

import "fmt"

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 | ~string
}

// Función genérica básica
func Imprimir[T any](valor T) {
	fmt.Printf("Valor: %v (Tipo: %T)\n", valor, valor)
}

// Función genérica con constraint comparable
func Contiene[T comparable](slice []T, elemento T) bool {
	for _, v := range slice {
		if v == elemento {
			return true
		}
	}
	return false
}

// Función genérica con constraint personalizada (usando golang.org/x/exp/constraints)
func Min[T Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Múltiples type parameters
func Pair[K comparable, V any](clave K, valor V) string {
	return fmt.Sprintf("%v: %v", clave, valor)
}

// Tipo genérico
type Par[K comparable, V any] struct {
	Clave K
	Valor V
}

func main() {
	// Type inference
	Imprimir(42)
	Imprimir("hola")
	Imprimir(3.14)

	// Con tipo explícito
	Imprimir[int](100)

	// Contiene (comparable)
	numeros := []int{1, 2, 3, 4, 5}
	fmt.Println("Contiene 3?", Contiene(numeros, 3))
	fmt.Println("Contiene 10?", Contiene(numeros, 10))

	palabras := []string{"go", "rust", "java"}
	fmt.Println("Contiene 'go'?", Contiene(palabras, "go"))

	// Min/Max (Ordered)
	fmt.Println("Min(10, 20):", Min(10, 20))
	fmt.Println("Max(3.14, 2.71):", Max(3.14, 2.71))
	fmt.Println("Min('a', 'z'):", Min("a", "z"))

	// Pair
	fmt.Println(Pair("nombre", "Alice"))

	// Tipo genérico Par
	p := Par[string, int]{Clave: "edad", Valor: 30}
	fmt.Printf("Par: %+v\n", p)
}
