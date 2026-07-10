package main

import "fmt"

func main() {
	// Crear mapas
	edades := make(map[string]int)
	edades["Alice"] = 30
	edades["Bob"] = 25
	edades["Charlie"] = 35

	// Map literal
	capitales := map[string]string{
		"España":  "Madrid",
		"Francia": "París",
		"Italia":  "Roma",
	}

	fmt.Println("Mapa de edades:", edades)
	fmt.Println("Capitales:", capitales)

	// Operaciones: agregar y actualizar
	edades["Diana"] = 28
	edades["Alice"] = 31
	fmt.Println("Después de agregar/actualizar:", edades)

	// Comprobación de existencia
	valor, existe := edades["Eve"]
	if existe {
		fmt.Println("Eve tiene", valor, "años")
	} else {
		fmt.Println("Eve no está en el mapa")
	}

	// delete
	delete(edades, "Bob")
	fmt.Println("Después de eliminar a Bob:", edades)

	// Recorrer mapas
	for nombre, edad := range edades {
		fmt.Printf("%s tiene %d años\n", nombre, edad)
	}

	// Longitud del mapa
	fmt.Println("Número de personas:", len(edades))

	// Comparación de existencia (comma ok idiom)
	if capital, ok := capitales["Portugal"]; ok {
		fmt.Println("Capital de Portugal:", capital)
	} else {
		fmt.Println("Portugal no está en el mapa")
	}
}
