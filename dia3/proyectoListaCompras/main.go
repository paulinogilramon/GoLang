package main

import (
	"fmt"
	"strings"
)

func main() {
	var lista []string

	for {
		fmt.Println("\n=== GESTOR DE LISTA DE COMPRAS ===")
		fmt.Println("1. Agregar artículo")
		fmt.Println("2. Ver lista")
		fmt.Println("3. Eliminar artículo")
		fmt.Println("4. Vaciar lista")
		fmt.Println("5. Buscar artículo")
		fmt.Println("0. Salir")
		fmt.Print("Elige: ")

		var opcion int
		fmt.Scan(&opcion)

		switch opcion {
		case 1:
			fmt.Print("Artículo a agregar: ")
			var item string
			fmt.Scan(&item)

			existe := false
			for _, a := range lista {
				if strings.EqualFold(a, item) {
					existe = true
					break
				}
			}
			if existe {
				fmt.Printf("'%s' ya está en la lista\n", item)
			} else {
				lista = append(lista, item)
				fmt.Printf("'%s' agregado (total: %d)\n", item, len(lista))
			}

		case 2:
			if len(lista) == 0 {
				fmt.Println("La lista está vacía")
			} else {
				fmt.Println("\n--- LISTA DE COMPRAS ---")
				for i, item := range lista {
					fmt.Printf("%d. %s\n", i+1, item)
				}
				fmt.Printf("Total: %d artículos\n", len(lista))
			}

		case 3:
			if len(lista) == 0 {
				fmt.Println("La lista está vacía")
				continue
			}
			fmt.Println("\n--- LISTA ACTUAL ---")
			for i, item := range lista {
				fmt.Printf("%d. %s\n", i+1, item)
			}
			fmt.Print("Número a eliminar: ")
			var idx int
			fmt.Scan(&idx)
			if idx < 1 || idx > len(lista) {
				fmt.Println("Número inválido")
			} else {
				eliminado := lista[idx-1]
				lista = append(lista[:idx-1], lista[idx:]...)
				fmt.Printf("'%s' eliminado\n", eliminado)
			}

		case 4:
			lista = nil
			fmt.Println("Lista vaciada")

		case 5:
			fmt.Print("Artículo a buscar: ")
			var busqueda string
			fmt.Scan(&busqueda)
			encontrados := []string{}
			for _, item := range lista {
				if strings.Contains(strings.ToLower(item), strings.ToLower(busqueda)) {
					encontrados = append(encontrados, item)
				}
			}
			if len(encontrados) == 0 {
				fmt.Println("No se encontraron coincidencias")
			} else {
				fmt.Println("Coincidencias:")
				for _, item := range encontrados {
					fmt.Println(" -", item)
				}
			}

		case 0:
			fmt.Println("¡Hasta luego!")
			return

		default:
			fmt.Println("Opción no válida")
		}
	}
}
