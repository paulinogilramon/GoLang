package main

import "fmt"

func main() {
	// 1. switch con valor
	dia := "martes"
	switch dia {
	case "lunes":
		fmt.Println("Empieza la semana")
	case "martes":
		fmt.Println("Segundo día")
	case "viernes":
		fmt.Println("Último día laboral")
	default:
		fmt.Println("Otro día cualquiera")
	}

	// 2. múltiples valores por case
	letra := "a"
	switch letra {
	case "a", "e", "i", "o", "u":
		fmt.Println("Es vocal")
	case "b", "c", "d":
		fmt.Println("Es consonante")
	default:
		fmt.Println("Otro carácter")
	}

	// 3. switch sin expresión (como if-else if)
	nota := 85
	switch {
	case nota >= 90:
		fmt.Println("Sobresaliente")
	case nota >= 70:
		fmt.Println("Notable")
	case nota >= 50:
		fmt.Println("Aprobado")
	default:
		fmt.Println("Reprobado")
	}

	// 4. fallthrough (forzar entrada al siguiente case)
	numero := 1
	switch numero {
	case 1:
		fmt.Println("Uno")
		fallthrough
	case 2:
		fmt.Println("Dos (por fallthrough)")
	default:
		fmt.Println("Otro")
	}

	fmt.Scanln()
}
