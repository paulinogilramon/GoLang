package main

import "fmt"

func main() {
	//1.Boolean
	var verdadero bool = true
	var falso bool

	//2.String
	var saludo string = "Hola Go"
	var vacio string

	//3.Enteros (con y sin signo, distintos tamaños)
	var (
		entero   int   = -42
		positivo uint  = 42
		pequeno  int8  = 127
		grande   int64 = 1_000_000_000
		unByte   byte  = 255 //alias de uint8
		unaRuna  rune  = 'Ñ' //alias de int32 (unicode)
	)

	//4.Flotantes
	var flotante32 float32 = 3.14159
	var flotante64 float64 = 2.718281828459045

	//5.Complejos
	var complejo complex64 = 1 + 2i
	var complejo128 complex128 = 3 + 4i

	//impresión con %T (tipo) y %v (valor)
	fmt.Println("=== Tipos Basicos ===")
	fmt.Printf("bool: %T -> %v\n", verdadero, verdadero)
	fmt.Printf("bool zero: %T -> %v\n", falso, falso)
	fmt.Printf("string: %T -> %q\n", saludo, saludo)
	fmt.Printf("string zero: %T -> %q\n", vacio, vacio)
	fmt.Printf("int: %T -> %d\n", entero, entero)
	fmt.Printf("uint: %T -> %d\n", positivo, positivo)
	fmt.Printf("int8: %T -> %d\n", pequeno, pequeno)
	fmt.Printf("int64: %T -> %d\n", grande, grande)
	fmt.Printf("byte: %T -> %d\n", unByte, unByte)
	fmt.Printf("rune: %T -> %c (%U)\n", unaRuna, unaRuna, unaRuna)
	fmt.Printf("float32: %T -> %f\n", flotante32, flotante32)
	fmt.Printf("float64: %T -> %f\n", flotante64, flotante64)
	fmt.Printf("complex64: %T -> %v\n", complejo, complejo)
	fmt.Printf("complex128: %T -> %v\n", complejo128, complejo128)

	//6.conversión explicita de tipos
	var (
		entero64 int64   = 100
		flotante float64 = float64(entero64)
		deVuelta int     = int(flotante)
	)
	fmt.Println("\n===Conversión explícita===")
	fmt.Printf("int64 -> float64: %f\n", flotante)
	fmt.Printf("float64 -> int: %d\n", deVuelta)

	fmt.Scanln()
}
