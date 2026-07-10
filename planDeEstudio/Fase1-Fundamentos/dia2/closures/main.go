package main

import "fmt"

func main() {
	// 1. Closure básico: función anónima asignada a variable
	suma := func(a, b int) int {
		return a + b
	}
	fmt.Println("1. Closure básico:", suma(3, 4))

	// 2. Closure que captura variables del entorno
	mensaje := "Hola"
	saludar := func(nombre string) {
		fmt.Println(mensaje, nombre)
	}
	saludar("Carlos")
	mensaje = "Adiós"
	saludar("Carlos")

	// 3. Closure como retorno de función (fábrica de funciones)
	incrementador := crearIncrementador(5)
	fmt.Println("3. Incremento:", incrementador())
	fmt.Println("   Incremento:", incrementador())
	fmt.Println("   Incremento:", incrementador())

	// 4. Closure para estado privado (contador)
	contador1 := nuevoContador()
	contador2 := nuevoContador()
	fmt.Println("4. Contador 1:", contador1())
	fmt.Println("   Contador 1:", contador1())
	fmt.Println("   Contador 2:", contador2())

	// 5. Closure con filtro (patrón funcional)
	numeros := []int{1, 2, 3, 4, 5, 6}
	pares := filtrar(numeros, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println("5. Pares:", pares)

	mayoresA3 := filtrar(numeros, func(n int) bool {
		return n > 3
	})
	fmt.Println("   Mayores a 3:", mayoresA3)

	// 6. Closure para operación diferida (callback)
	ejecutarCallback(func() {
		fmt.Println("6. Callback ejecutado!")
	})

	fmt.Scanln()
}

// crearIncrementador retorna un closure que incrementa y devuelve un valor
func crearIncrementador(inicial int) func() int {
	acumulador := inicial
	return func() int {
		acumulador++
		return acumulador
	}
}

// nuevoContador retorna un closure que mantiene su propio contador
func nuevoContador() func() int {
	contador := 0
	return func() int {
		contador++
		return contador
	}
}

// filtrar aplica un predicado a un slice y devuelve los elementos que lo cumplen
func filtrar(numeros []int, predicado func(int) bool) []int {
	var resultado []int
	for _, n := range numeros {
		if predicado(n) {
			resultado = append(resultado, n)
		}
	}
	return resultado
}

// ejecutarCallback recibe una función y la ejecuta
func ejecutarCallback(callback func()) {
	fmt.Println("   Preparando...")
	callback()
}
