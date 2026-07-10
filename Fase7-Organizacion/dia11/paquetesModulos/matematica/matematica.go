package matematica

import "math"

// Suma retorna la suma de dos números
func Suma(a, b float64) float64 {
	return a + b
}

// Resta retorna a - b
func Resta(a, b float64) float64 {
	return a - b
}

// Multiplica retorna a * b
func Multiplica(a, b float64) float64 {
	return a * b
}

// Divide retorna a / b. Retorna 0 si b es 0.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrorDivisionCero()
	}
	return a / b, nil
}

// Potencia retorna a^b
func Potencia(a, b float64) float64 {
	return math.Pow(a, b)
}

// Factorial retorna n! para enteros no negativos
func Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1)
}

// Fibonacci retorna el n-ésimo número de Fibonacci
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}
