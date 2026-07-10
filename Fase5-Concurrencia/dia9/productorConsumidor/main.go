package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Tarea struct {
	ID   int
	Data string
}

func Productor(id int, tareas chan<- Tarea, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		t := Tarea{
			ID:   i + id*100,
			Data: fmt.Sprintf("Tarea del productor %d (item %d)", id, i),
		}
		fmt.Printf("Productor %d -> produce: %s\n", id, t.Data)
		tareas <- t
	}
}

func Consumidor(id int, tareas <-chan Tarea, wg *sync.WaitGroup) {
	defer wg.Done()
	for tarea := range tareas {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		fmt.Printf("Consumidor %d <- procesa [%d]: %s\n", id, tarea.ID, tarea.Data)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	tareas := make(chan Tarea, 10)
	var wg sync.WaitGroup

	// Iniciar productores
	numProductores := 3
	for i := 1; i <= numProductores; i++ {
		wg.Add(1)
		go Productor(i, tareas, &wg)
	}

	// Iniciar consumidores
	var wgConsumidores sync.WaitGroup
	numConsumidores := 2
	for i := 1; i <= numConsumidores; i++ {
		wgConsumidores.Add(1)
		go Consumidor(i, tareas, &wgConsumidores)
	}

	// Esperar a que todos los productores terminen
	wg.Wait()
	close(tareas) // cerrar canal para que consumidores terminen

	// Esperar a que consumidores terminen de procesar
	wgConsumidores.Wait()
	fmt.Println("\nSistema productor-consumidor finalizado.")
}
