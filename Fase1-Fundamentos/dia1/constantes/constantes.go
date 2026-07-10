package main

import "fmt"

func main() {
	//1. Constante simple NO tipada (Go infiere el tipo)
	const Pi = 3.14159

	//2. Constante simple tipada
	const EdadMaxima int = 120

	//3. Multiples constantes en bloque
	const (
		Lunes  = 1
		Martes = 2
	)

	//4. iota - enumerador automatico
	const (
		Domingo = iota
		LunesI
		MartesI
	)

	//5. iota con salto (usando "_")
	const (
		Uno  = iota + 1 //empeiza en 1
		_               //se salta el 2
		Tres            //3
	)

	//Imprimir todo
	fmt.Println("Pi:", Pi)
	fmt.Println("EdadMaxima:", EdadMaxima)
	fmt.Println("Lunes:", Lunes, "Martes:", Martes)
	fmt.Println("Domingo:", Domingo, "LunesI:", LunesI, "MartesI:", MartesI)

	//Dato curioso: Go No deja usar := con const
	//const x := 5 //eso da error de compilacion

	fmt.Scanln()
}
