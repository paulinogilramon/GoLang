package main

import (
	"fmt"
	"time"
)

func main() {
	// Channel sin buffer (bloqueante)
	mensajes := make(chan string)
	go func() {
		mensajes <- "hola"
	}()
	msg := <-mensajes
	fmt.Println("Channel sin buffer:", msg)

	// Channel con buffer
	canalBuf := make(chan int, 3)
	canalBuf <- 1
	canalBuf <- 2
	canalBuf <- 3
	fmt.Println("Channel con buffer:")
	fmt.Println(" ", <-canalBuf)
	fmt.Println(" ", <-canalBuf)
	fmt.Println(" ", <-canalBuf)

	// Channel direccional
	soloEnvio := make(chan<- string) // solo escritura
	soloLectura := make(<-chan string) // solo lectura
	_ = soloEnvio
	_ = soloLectura

	// Range sobre channel
	numeros := make(chan int, 5)
	for i := 0; i < 5; i++ {
		numeros <- i * 10
	}
	close(numeros)
	for n := range numeros {
		fmt.Println("Range channel:", n)
	}

	// select con múltiples channels
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "mensaje del canal 1"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "mensaje del canal 2"
	}()

	select {
	case msg1 := <-c1:
		fmt.Println("Recibido de c1:", msg1)
	case msg2 := <-c2:
		fmt.Println("Recibido de c2:", msg2)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout")
	}

	// select con default (non-blocking)
	select {
	case msg := <-c1:
		fmt.Println("Recibido:", msg)
	default:
		fmt.Println("No hay mensajes disponibles (non-blocking)")
	}
}
