package main

import (
	"errors"
	"fmt"
)

// Error personalizado con "sentinel error"
var ErrSalarioNegativo = errors.New("el salario no puede ser negativo")
var ErrSalarioMuyBajo = errors.New("el salario mínimo es 1000")

// Error personalizado como struct
type ErrorValidacion struct {
	Campo   string
	Mensaje string
}

func (e *ErrorValidacion) Error() string {
	return fmt.Sprintf("error en campo '%s': %s", e.Campo, e.Mensaje)
}

// Función que retorna error
func CalcularSalario(horas int, tarifa float64) (float64, error) {
	if horas < 0 {
		return 0, errors.New("las horas no pueden ser negativas")
	}
	if tarifa < 0 {
		return 0, ErrSalarioNegativo
	}
	salario := float64(horas) * tarifa
	if salario < 1000 {
		return 0, ErrSalarioMuyBajo
	}
	return salario, nil
}

// Función con fmt.Errorf
func ValidarEdad(edad int) error {
	if edad < 0 {
		return fmt.Errorf("edad inválida: %d (debe ser positiva)", edad)
	}
	if edad > 150 {
		return fmt.Errorf("edad inválida: %d (máximo 150)", edad)
	}
	return nil
}

// defer, panic y recover
func EjemploDeferPanicRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recuperado de panic:", r)
		}
	}()

	fmt.Println("Ejecutando función...")
	defer fmt.Println("Esto se ejecuta al final (defer)")

	panic("¡algo salió muy mal!")
	// Esto nunca se ejecuta
	fmt.Println("Esto no se imprime")
}

func main() {
	// Manejo básico de errores
	salario, err := CalcularSalario(40, 50)
	if err != nil {
		if errors.Is(err, ErrSalarioNegativo) {
			fmt.Println("Error de salario negativo:", err)
		} else {
			fmt.Println("Error:", err)
		}
	} else {
		fmt.Println("Salario calculado:", salario)
	}

	// Validación con error personalizado
	err = ValidarEdad(-5)
	if err != nil {
		fmt.Println(err)
	}

	// Error como struct
	errValidacion := &ErrorValidacion{Campo: "email", Mensaje: "formato inválido"}
	fmt.Println(errValidacion)

	// Errors.Is y errors.As
	_, err = CalcularSalario(-10, 50)
	if errors.Is(err, ErrSalarioMuyBajo) {
		fmt.Println("El salario es muy bajo")
	} else {
		fmt.Println("Otro error:", err)
	}

	// defer, panic, recover
	fmt.Println("\n--- defer/panic/recover ---")
	EjemploDeferPanicRecover()
	fmt.Println("El programa continúa después del panic.")

	// defer típico: cerrar recursos
	simularArchivo := func(nombre string) {
		defer fmt.Printf("Cerrando archivo: %s\n", nombre)
		fmt.Printf("Abriendo archivo: %s\n", nombre)
	}
	simularArchivo("datos.txt")
}
