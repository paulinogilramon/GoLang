package main

import (
	"fmt"
	"time"
)

type ResultadoProcesamiento struct {
	TareaID int
	Exitoso bool
	Error   string
	Duracion time.Duration
}

func ProcesadorTareas(gestor *GestorTareas, tareas <-chan int, resultados chan<- ResultadoProcesamiento, done chan<- bool) {
	for id := range tareas {
		inicio := time.Now()
		time.Sleep(500 * time.Millisecond)

		_, err := gestor.Obtener(id)
		if err != nil {
			resultados <- ResultadoProcesamiento{
				TareaID: id,
				Exitoso: false,
				Error:   err.Error(),
				Duracion: time.Since(inicio),
			}
			continue
		}

		_, err = gestor.Actualizar(id, "", "", "", EnProgres)
		if err != nil {
			resultados <- ResultadoProcesamiento{
				TareaID: id,
				Exitoso: false,
				Error:   err.Error(),
				Duracion: time.Since(inicio),
			}
			continue
		}

		time.Sleep(500 * time.Millisecond)

		_, err = gestor.Actualizar(id, "", "", "", Completada)
		if err != nil {
			resultados <- ResultadoProcesamiento{
				TareaID: id,
				Exitoso: false,
				Error:   err.Error(),
				Duracion: time.Since(inicio),
			}
			continue
		}

		resultados <- ResultadoProcesamiento{
			TareaID:  id,
			Exitoso:  true,
			Duracion: time.Since(inicio),
		}
	}
	done <- true
}

func DemostrarConcurrencia(gestor *GestorTareas) {
	fmt.Println("\n=== Demostración de Concurrencia ===")

	for i := 1; i <= 5; i++ {
		gestor.Agregar(
			fmt.Sprintf("Tarea concurrente %d", i),
			fmt.Sprintf("Descripción de la tarea %d", i),
			"concurrente",
		)
	}

	tareas := make(chan int, 10)
	resultados := make(chan ResultadoProcesamiento, 10)
	done := make(chan bool, 3)

	numWorkers := 3
	for w := 1; w <= numWorkers; w++ {
		go ProcesadorTareas(gestor, tareas, resultados, done)
	}

	go func() {
		for id := 1; id <= 5; id++ {
			tareas <- id
		}
		close(tareas)
	}()

	go func() {
		for i := 0; i < numWorkers; i++ {
			<-done
		}
		close(resultados)
	}()

	for res := range resultados {
		status := "✓"
		if !res.Exitoso {
			status = "✗"
		}
		fmt.Printf("  [%s] Tarea %d procesada en %v", status, res.TareaID, res.Duracion.Round(time.Millisecond))
		if !res.Exitoso {
			fmt.Printf(" - Error: %s", res.Error)
		}
		fmt.Println()
	}

	fmt.Println("Procesamiento concurrente finalizado.")
}
