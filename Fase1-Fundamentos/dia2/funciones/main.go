package main

import "fmt"

//1.Función simple sin parámetros ni retorno
func saludar() {
	fmt.Println("Hola desde la función saludar!")
}

//2.Función con parámetros
func saludarPersona(nombre string) {
	fmt.Printf("Hola %s!\n", nombre)
}

//3.Función con parámetros y retorno
func sumar(a int, b int) int {
	return a + b
}

//4.Función con parámetros del mismo tipo (sintaxis compacta)
func multiplicar(a, b, c int) int {
	return a * b * c
}

//5.Función con retorno nombrado
func dividir(dividiendo, divisor int) (resultado int) {
	resultado = dividiendo / divisor
	return //return "desnudo" devuelve resultado
}

func main() {
	saludar()
	saludarPersona("Carlos")

	s := sumar(10, 5)
	fmt.Println("Suma:", s)

	m := multiplicar(2, 3, 4)
	fmt.Println("Multiplicación:", m)

	d := dividir(20, 4)
	fmt.Println("División:", d)

	fmt.Scanln()
}
