package main

import "fmt"

type Estado string

const (
	Pendiente Estado = "pendiente"
	EnProgres Estado = "en_progreso"
	Completada Estado = "completada"
)

type Tarea struct {
	ID          int
	Titulo      string
	Descripcion string
	Estado      Estado
	Categoria   string
}

func (t Tarea) Mostrar() {
	icono := map[Estado]string{
		Pendiente:  "[ ]",
		EnProgres:  "[~]",
		Completada: "[✓]",
	}[t.Estado]
	fmt.Printf("%s #%d: %s (%s)\n  %s\n", icono, t.ID, t.Titulo, t.Categoria, t.Descripcion)
}

type Tareas []Tarea

func (ts Tareas) FiltrarPorEstado(estado Estado) Tareas {
	var filtradas Tareas
	for _, t := range ts {
		if t.Estado == estado {
			filtradas = append(filtradas, t)
		}
	}
	return filtradas
}

func (ts Tareas) Buscar(query string) Tareas {
	var resultados Tareas
	for _, t := range ts {
		if contiene(t.Titulo, query) || contiene(t.Descripcion, query) {
			resultados = append(resultados, t)
		}
	}
	return resultados
}

func contiene(s, sub string) bool {
	return len(s) >= len(sub) && containsStr(s, sub)
}

func containsStr(s, sub string) bool {
	for i := 0; i <= len(s)-len(sub); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}
