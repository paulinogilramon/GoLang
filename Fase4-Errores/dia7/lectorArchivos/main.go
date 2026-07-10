package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

var ErrArchivoNoEncontrado = errors.New("archivo no encontrado")

type ArchivoError struct {
	Nombre string
	Err    error
}

func (e *ArchivoError) Error() string {
	return fmt.Sprintf("error con archivo '%s': %v", e.Nombre, e.Err)
}

func (e *ArchivoError) Unwrap() error {
	return e.Err
}

func LeerArchivo(nombre string) ([]string, error) {
	f, err := os.Open(nombre)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("%w: %s", ErrArchivoNoEncontrado, nombre)
		}
		return nil, &ArchivoError{Nombre: nombre, Err: err}
	}
	defer f.Close()

	var lineas []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lineas = append(lineas, scanner.Text())
	}
	return lineas, scanner.Err()
}

func main() {
	// Crear un archivo de prueba temporal
	tmpFile := "datos_prueba.txt"
	contenido := []string{"linea 1: hola mundo", "linea 2: esto es go", "linea 3: manejo de archivos"}
	os.WriteFile(tmpFile, []byte(strings.Join(contenido, "\n")), 0644)
	defer os.Remove(tmpFile)

	// Leer archivo existente
	lineas, err := LeerArchivo(tmpFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Contenido del archivo:")
	for _, linea := range lineas {
		fmt.Println(" ", linea)
	}

	// Leer archivo inexistente (manejo robusto)
	fmt.Println("\nIntentando leer archivo inexistente:")
	_, err = LeerArchivo("no_existe.txt")
	if err != nil {
		if errors.Is(err, ErrArchivoNoEncontrado) {
			fmt.Println("Error controlado:", err)
		} else {
			fmt.Println("Error inesperado:", err)
		}
	}
}
