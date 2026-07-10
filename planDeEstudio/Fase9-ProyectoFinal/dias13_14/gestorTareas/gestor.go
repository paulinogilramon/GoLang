package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sort"
	"sync"
)

var (
	ErrTareaNoEncontrada = errors.New("tarea no encontrada")
	ErrEstadoInvalido    = errors.New("estado inválido")
)

type GestorTareas struct {
	mu       sync.RWMutex
	tareas   map[int]Tarea
	proximoID int
	archivo  string
}

func NuevoGestor(archivo string) *GestorTareas {
	return &GestorTareas{
		tareas:   make(map[int]Tarea),
		proximoID: 1,
		archivo:  archivo,
	}
}

func (g *GestorTareas) Agregar(titulo, descripcion, categoria string) Tarea {
	g.mu.Lock()
	defer g.mu.Unlock()

	t := Tarea{
		ID:          g.proximoID,
		Titulo:      titulo,
		Descripcion: descripcion,
		Estado:      Pendiente,
		Categoria:   categoria,
	}
	g.tareas[t.ID] = t
	g.proximoID++
	return t
}

func (g *GestorTareas) Obtener(id int) (Tarea, error) {
	g.mu.RLock()
	defer g.mu.RUnlock()

	t, ok := g.tareas[id]
	if !ok {
		return Tarea{}, fmt.Errorf("%w: %d", ErrTareaNoEncontrada, id)
	}
	return t, nil
}

func (g *GestorTareas) Actualizar(id int, titulo, descripcion, categoria string, estado Estado) (Tarea, error) {
	g.mu.Lock()
	defer g.mu.Unlock()

	t, ok := g.tareas[id]
	if !ok {
		return Tarea{}, fmt.Errorf("%w: %d", ErrTareaNoEncontrada, id)
	}

	if titulo != "" {
		t.Titulo = titulo
	}
	if descripcion != "" {
		t.Descripcion = descripcion
	}
	if categoria != "" {
		t.Categoria = categoria
	}
	if estado != "" {
		switch estado {
		case Pendiente, EnProgres, Completada:
			t.Estado = estado
		default:
			return Tarea{}, fmt.Errorf("%w: %s", ErrEstadoInvalido, estado)
		}
	}

	g.tareas[id] = t
	return t, nil
}

func (g *GestorTareas) Eliminar(id int) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	if _, ok := g.tareas[id]; !ok {
		return fmt.Errorf("%w: %d", ErrTareaNoEncontrada, id)
	}
	delete(g.tareas, id)
	return nil
}

func (g *GestorTareas) Listar() Tareas {
	g.mu.RLock()
	defer g.mu.RUnlock()

	lista := make(Tareas, 0, len(g.tareas))
	for _, t := range g.tareas {
		lista = append(lista, t)
	}
	sort.Slice(lista, func(i, j int) bool {
		return lista[i].ID < lista[j].ID
	})
	return lista
}

func (g *GestorTareas) ContarPorEstado() map[Estado]int {
	g.mu.RLock()
	defer g.mu.RUnlock()

	conteo := map[Estado]int{
		Pendiente:  0,
		EnProgres:  0,
		Completada: 0,
	}
	for _, t := range g.tareas {
		conteo[t.Estado]++
	}
	return conteo
}

func (g *GestorTareas) GuardarJSON() error {
	g.mu.RLock()
	defer g.mu.RUnlock()

	data, err := json.MarshalIndent(g.tareas, "", "  ")
	if err != nil {
		return fmt.Errorf("error al serializar: %w", err)
	}
	return os.WriteFile(g.archivo, data, 0644)
}

func (g *GestorTareas) CargarJSON() error {
	g.mu.Lock()
	defer g.mu.Unlock()

	data, err := os.ReadFile(g.archivo)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("error al leer archivo: %w", err)
	}

	var tareas map[int]Tarea
	if err := json.Unmarshal(data, &tareas); err != nil {
		return fmt.Errorf("error al deserializar: %w", err)
	}

	g.tareas = tareas
	for id := range tareas {
		if id >= g.proximoID {
			g.proximoID = id + 1
		}
	}
	return nil
}
