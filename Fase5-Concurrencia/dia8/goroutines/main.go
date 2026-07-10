package main

import (
	"fmt"
	"sync"
	"time"
)

func tarea(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Goroutine %d: empezando\n", id)
	time.Sleep(time.Duration(id) * 500 * time.Millisecond)
	fmt.Printf("Goroutine %d: terminando\n", id)
}

func main() {
	// Goroutine básica
	go func() {
		fmt.Println("Hello desde una goroutine!")
	}()

	// sync.WaitGroup para esperar goroutines
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go tarea(i, &wg)
	}

	// Múltiples goroutines con closures
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Printf("Closure goroutine %d\n", n)
		}(i)
	}

	wg.Wait()
	fmt.Println("Todas las goroutines terminaron.")

	// Race condition (demostración)
	var contador int
	var wg2 sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 1000; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			mu.Lock()
			contador++
			mu.Unlock()
		}()
	}
	wg2.Wait()
	fmt.Println("Contador con mutex:", contador)
}
