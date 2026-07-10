package main

import (
	"fmt"
	"math"
)

const PI = 3.1416
const GRAVEDAD = 9.8

func menu() int {
	var opcion int
	fmt.Println("\n=== BIBLIOTECA MATEMÁTICA ===")
	fmt.Println("1. Operación con closure (suma/resta/multiplica)")
	fmt.Println("2. Secuencia personalizada")
	fmt.Println("3. Estadísticas básicas")
	fmt.Println("4. Resolver ecuación lineal (ax + b = 0)")
	fmt.Println("5. Calcular velocidad final (MRUV)")
	fmt.Println("6. Calculadora con switch")
	fmt.Println("7. Salir")
	fmt.Print("Elige: ")
	fmt.Scan(&opcion)
	return opcion
}

func crearOperador(operador string) func(float64, float64) (float64, error) {
	switch operador {
	case "+":
		return func(a, b float64) (float64, error) {
			return a + b, nil
		}
	case "-":
		return func(a, b float64) (float64, error) {
			return a - b, nil
		}
	case "*":
		return func(a, b float64) (float64, error) {
			return a * b, nil
		}
	case "/":
		return func(a, b float64) (float64, error) {
			if b == 0 {
				return 0, fmt.Errorf("división entre cero")
			}
			return a / b, nil
		}
	default:
		return func(a, b float64) (float64, error) {
			return 0, fmt.Errorf("operador '%s' no soportado", operador)
		}
	}
}

func crearSecuencia(inicio, paso float64) func() float64 {
	actual := inicio
	return func() float64 {
		valor := actual
		actual += paso
		return valor
	}
}

func estadisticas(numeros ...float64) (suma, promedio, min, max float64) {
	if len(numeros) == 0 {
		return 0, 0, 0, 0
	}
	suma = 0
	min = numeros[0]
	max = numeros[0]
	for _, n := range numeros {
		suma += n
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	promedio = suma / float64(len(numeros))
	return
}

func ecuacionLineal(a, b float64) (float64, error) {
	if a == 0 {
		return 0, fmt.Errorf("'a' no puede ser 0")
	}
	return -b / a, nil
}

func velocidadFinal(velocidadInicial, aceleracion, tiempo float64) (float64, error) {
	if tiempo < 0 {
		return 0, fmt.Errorf("el tiempo no puede ser negativo")
	}
	return velocidadInicial + aceleracion*tiempo, nil
}

func main() {
	for {
		opcion := menu()
		switch opcion {
		case 1:
			var a, b float64
			var op string
			fmt.Print("Ingresa operación (ej: 3 + 5): ")
			fmt.Scan(&a, &op, &b)
			operacion := crearOperador(op)
			resultado, err := operacion(a, b)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("%.2f %s %.2f = %.2f\n", a, op, b, resultado)
			}

		case 2:
			var inicio, paso float64
			var cantidad int
			fmt.Print("Inicio, paso, cantidad (ej: 0 2 5): ")
			fmt.Scan(&inicio, &paso, &cantidad)
			secuencia := crearSecuencia(inicio, paso)
			fmt.Print("Secuencia:")
			for i := 0; i < cantidad; i++ {
				fmt.Printf(" %.0f", secuencia())
			}
			fmt.Println()

		case 3:
			var n int
			fmt.Print("¿Cuántos números?: ")
			fmt.Scan(&n)
			if n <= 1 {
				fmt.Println("Necesitas al menos 2 números")
				continue
			}
			numeros := make([]float64, n)
			for i := 0; i < n; i++ {
				fmt.Printf("Número %d: ", i+1)
				fmt.Scan(&numeros[i])
			}
			suma, promedio, min, max := estadisticas(numeros...)
			fmt.Printf("Suma: %.2f\n", suma)
			fmt.Printf("Promedio: %.2f\n", promedio)
			fmt.Printf("Mínimo: %.2f\n", min)
			fmt.Printf("Máximo: %.2f\n", max)

		case 4:
			var a, b float64
			fmt.Print("Ingresa a y b (ax + b = 0): ")
			fmt.Scan(&a, &b)
			resultado, err := ecuacionLineal(a, b)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("x = %.2f\n", resultado)
			}

		case 5:
			var v0, a, t float64
			fmt.Print("Velocidad inicial, aceleración, tiempo: ")
			fmt.Scan(&v0, &a, &t)
			vf, err := velocidadFinal(v0, a, t)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Velocidad final: %.2f m/s\n", vf)
				fmt.Printf("(usando constante GRAVEDAD = %.1f m/s²)\n", GRAVEDAD)
				if t >= 1 {
					distancia := v0*t + 0.5*a*t*t
					fmt.Printf("Distancia recorrida: %.2f m\n", distancia)
				}
			}

		case 6:
			var a, b float64
			var operador string
			fmt.Print("Calculadora switch (ej: 10 / 3): ")
			fmt.Scan(&a, &operador, &b)
			var resultado float64
			switch operador {
			case "+":
				resultado = a + b
			case "-":
				resultado = a - b
			case "*":
				resultado = a * b
			case "/":
				if b == 0 {
					fmt.Println("Error: división entre cero")
					continue
				}
				resultado = a / b
			case "^":
				resultado = math.Pow(a, b)
			default:
				fmt.Println("Operador no válido")
				continue
			}
			fmt.Printf("%.2f %s %.2f = %.2f\n", a, operador, b, resultado)

		case 7:
			fmt.Println("¡Hasta luego!")
			return

		default:
			fmt.Println("Opción inválida")
		}
	}
}
