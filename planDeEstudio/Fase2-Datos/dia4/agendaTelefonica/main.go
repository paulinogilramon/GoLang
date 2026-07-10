package main

import (
	"fmt"
	"strings"
)

type Contacto struct {
	Nombre  string
	Telefono string
	Email   string
}

func main() {
	agenda := make(map[string]Contacto)

	for {
		fmt.Println("\n=== AGENDA TELEFÓNICA ===")
		fmt.Println("1. Agregar contacto")
		fmt.Println("2. Buscar contacto")
		fmt.Println("3. Eliminar contacto")
		fmt.Println("4. Listar todos")
		fmt.Println("5. Salir")
		fmt.Print("Elige una opción: ")

		var opcion int
		fmt.Scan(&opcion)

		switch opcion {
		case 1:
			var nombre, telefono, email string
			fmt.Print("Nombre: ")
			fmt.Scan(&nombre)
			fmt.Print("Teléfono: ")
			fmt.Scan(&telefono)
			fmt.Print("Email: ")
			fmt.Scan(&email)
			agenda[strings.ToLower(nombre)] = Contacto{nombre, telefono, email}
			fmt.Println("Contacto agregado.")

		case 2:
			var nombre string
			fmt.Print("Nombre a buscar: ")
			fmt.Scan(&nombre)
			if c, ok := agenda[strings.ToLower(nombre)]; ok {
				fmt.Printf("Nombre: %s | Tel: %s | Email: %s\n", c.Nombre, c.Telefono, c.Email)
			} else {
				fmt.Println("Contacto no encontrado.")
			}

		case 3:
			var nombre string
			fmt.Print("Nombre a eliminar: ")
			fmt.Scan(&nombre)
			if _, ok := agenda[strings.ToLower(nombre)]; ok {
				delete(agenda, strings.ToLower(nombre))
				fmt.Println("Contacto eliminado.")
			} else {
				fmt.Println("Contacto no encontrado.")
			}

		case 4:
			if len(agenda) == 0 {
				fmt.Println("Agenda vacía.")
			} else {
				for _, c := range agenda {
					fmt.Printf("Nombre: %s | Tel: %s | Email: %s\n", c.Nombre, c.Telefono, c.Email)
				}
			}

		case 5:
			fmt.Println("¡Hasta luego!")
			return

		default:
			fmt.Println("Opción inválida.")
		}
	}
}
