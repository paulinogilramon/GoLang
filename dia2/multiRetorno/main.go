package main

import "fmt"

// 1.Función con múltiples valores de retorno
func dividirConResto(a, b int) (int, int) {
	cociente := a / b
	resto := a % b
	return cociente, resto
}

// 2.Retornos nombrados
func operaciones(a, b int) (suma, resta, producto int) {
	suma = a + b
	resta = a - b
	producto = a * b
	return
}

// 3.Patrón común: retornar valor + error
func dividirSegura(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("no se puede dividir por cero")
	}
	return a / b, nil
}

// 4.Función con diferentes tipos de retorno
func estadisticas(numeros ...int) (float64, int, int) {
	if len(numeros) == 0 {
		return 0, 0, 0
	}
	suma := 0
	min, max := numeros[0], numeros[0]
	for _, n := range numeros {
		suma += n
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	promedio := float64(suma) / float64(len(numeros))
	return promedio, min, max
}

func main() {
	//Recibir múltiples valores
	cociente, resto := dividirConResto(17, 5)
	fmt.Printf("17 / 5 = %d, resto = %d\n", cociente, resto)

	//Ignorar valores con _
	soloCociente, _ := dividirConResto(20, 3)
	fmt.Println("Solo cociente:", soloCociente)

	//Retornos nombrados
	s, r, p := operaciones(10, 3)
	fmt.Printf("Suma: %d, Resta: %d, Producto: %d\n", s, r, p)

	//Patrón valor + error
	resultado, err := dividirSegura(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("División segura:", resultado)
	}

	//División por cero
	_, err2 := dividirSegura(10, 0)
	if err2 != nil {
		fmt.Println("Error:", err2)
	}

	//Retorno con tipos mixtos
	prom, mn, mx := estadisticas(4, 7, 2, 9, 5)
	fmt.Printf("Promedio: %.2f, Mínimo: %d, Máximo: %d\n", prom, mn, mx)

	fmt.Scanln()
}
