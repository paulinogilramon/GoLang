package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	defer fmt.Println("Gestor de Tareas cerrado.")

	gestor := NuevoGestor("tareas.json")
	if err := gestor.CargarJSON(); err != nil {
		fmt.Println("Aviso: no se pudieron cargar tareas previas:", err)
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n═══════════════════════════════════")
		fmt.Println("        GESTOR DE TAREAS")
		fmt.Println("═══════════════════════════════════")
		fmt.Println(" 1. Agregar tarea")
		fmt.Println(" 2. Listar tareas")
		fmt.Println(" 3. Ver tarea")
		fmt.Println(" 4. Actualizar tarea")
		fmt.Println(" 5. Eliminar tarea")
		fmt.Println(" 6. Buscar tareas")
		fmt.Println(" 7. Filtrar por estado")
		fmt.Println(" 8. Estadísticas")
		fmt.Println(" 9. Demo concurrencia")
		fmt.Println("10. Guardar y salir")
		fmt.Print("Elige: ")

		scanner.Scan()
		opcion := strings.TrimSpace(scanner.Text())

		switch opcion {
		case "1":
			agregarTarea(gestor, scanner)
		case "2":
			listarTareas(gestor)
		case "3":
			verTarea(gestor, scanner)
		case "4":
			actualizarTarea(gestor, scanner)
		case "5":
			eliminarTarea(gestor, scanner)
		case "6":
			buscarTareas(gestor, scanner)
		case "7":
			filtrarTareas(gestor, scanner)
		case "8":
			mostrarEstadisticas(gestor)
		case "9":
			DemostrarConcurrencia(gestor)
		case "10":
			if err := gestor.GuardarJSON(); err != nil {
				fmt.Println("Error al guardar:", err)
				return
			}
			fmt.Println("Tareas guardadas. ¡Hasta luego!")
			return
		default:
			fmt.Println("Opción inválida.")
		}
	}
}

func leerCampo(scanner *bufio.Scanner, prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func agregarTarea(g *GestorTareas, s *bufio.Scanner) {
	titulo := leerCampo(s, "Título: ")
	if titulo == "" {
		fmt.Println("El título no puede estar vacío.")
		return
	}
	desc := leerCampo(s, "Descripción: ")
	cat := leerCampo(s, "Categoría: ")
	t := g.Agregar(titulo, desc, cat)
	fmt.Printf("Tarea #%d creada.\n", t.ID)
}

func listarTareas(g *GestorTareas) {
	tareas := g.Listar()
	if len(tareas) == 0 {
		fmt.Println("No hay tareas.")
		return
	}
	for _, t := range tareas {
		t.Mostrar()
	}
}

func verTarea(g *GestorTareas, s *bufio.Scanner) {
	id := leerID(s)
	t, err := g.Obtener(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Mostrar()
}

func actualizarTarea(g *GestorTareas, s *bufio.Scanner) {
	id := leerID(s)

	titulo := leerCampo(s, "Nuevo título (vacío = sin cambio): ")
	desc := leerCampo(s, "Nueva descripción (vacío = sin cambio): ")
	cat := leerCampo(s, "Nueva categoría (vacío = sin cambio): ")

	estadoStr := leerCampo(s, "Estado (pendiente/en_progreso/completada/vacío = sin cambio): ")
	var estado Estado
	switch estadoStr {
	case "pendiente":
		estado = Pendiente
	case "en_progreso":
		estado = EnProgres
	case "completada":
		estado = Completada
	case "":
	default:
		fmt.Println("Estado inválido.")
		return
	}

	t, err := g.Actualizar(id, titulo, desc, cat, estado)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Tarea actualizada:")
	t.Mostrar()
}

func eliminarTarea(g *GestorTareas, s *bufio.Scanner) {
	id := leerID(s)
	if err := g.Eliminar(id); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Tarea #%d eliminada.\n", id)
}

func buscarTareas(g *GestorTareas, s *bufio.Scanner) {
	query := leerCampo(s, "Buscar: ")
	resultados := Tareas(g.Listar()).Buscar(query)
	if len(resultados) == 0 {
		fmt.Println("Sin resultados.")
		return
	}
	for _, t := range resultados {
		t.Mostrar()
	}
}

func filtrarTareas(g *GestorTareas, s *bufio.Scanner) {
	fmt.Println("Estados: pendiente, en_progreso, completada")
	estadoStr := leerCampo(s, "Filtrar por estado: ")
	var estado Estado
	switch estadoStr {
	case "pendiente":
		estado = Pendiente
	case "en_progreso":
		estado = EnProgres
	case "completada":
		estado = Completada
	default:
		fmt.Println("Estado inválido.")
		return
	}
	filtradas := Tareas(g.Listar()).FiltrarPorEstado(estado)
	for _, t := range filtradas {
		t.Mostrar()
	}
}

func mostrarEstadisticas(g *GestorTareas) {
	conteo := g.ContarPorEstado()
	total := 0
	for _, c := range conteo {
		total += c
	}
	fmt.Printf("Total: %d tareas\n", total)
	for est, count := range conteo {
		fmt.Printf("  %s: %d\n", est, count)
	}
}

func leerID(s *bufio.Scanner) int {
	s.Scan()
	var id int
	fmt.Sscanf(strings.TrimSpace(s.Text()), "%d", &id)
	return id
}
