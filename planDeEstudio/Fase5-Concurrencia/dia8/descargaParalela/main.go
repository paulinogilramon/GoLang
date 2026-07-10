package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type ArchivoDescarga struct {
	Nombre      string
	URL         string
	Bytes       int
	Duracion    time.Duration
}

func Descargar(archivo ArchivoDescarga, wg *sync.WaitGroup, resultados chan<- string) {
	defer wg.Done()

	duracion := time.Duration(rand.Intn(2000)+500) * time.Millisecond
	time.Sleep(duracion)

	bytes := rand.Intn(10000) + 1000
	resultado := fmt.Sprintf("✓ %s descargado (%d bytes en %v)", archivo.Nombre, bytes, duracion.Round(time.Millisecond))
	resultados <- resultado
}

func main() {
	rand.Seed(time.Now().UnixNano())

	archivos := []ArchivoDescarga{
		{Nombre: "foto1.jpg", URL: "http://ejemplo.com/foto1.jpg"},
		{Nombre: "foto2.jpg", URL: "http://ejemplo.com/foto2.jpg"},
		{Nombre: "documento.pdf", URL: "http://ejemplo.com/doc.pdf"},
		{Nombre: "video.mp4", URL: "http://ejemplo.com/video.mp4"},
		{Nombre: "musica.mp3", URL: "http://ejemplo.com/musica.mp3"},
		{Nombre: "datos.zip", URL: "http://ejemplo.com/datos.zip"},
	}

	fmt.Println("=== Descargando archivos en paralelo ===\n")

	var wg sync.WaitGroup
	resultados := make(chan string, len(archivos))

	inicio := time.Now()

	for _, archivo := range archivos {
		wg.Add(1)
		go Descargar(archivo, &wg, resultados)
	}

	wg.Wait()
	close(resultados)

	for resultado := range resultados {
		fmt.Println(resultado)
	}

	fmt.Printf("\nTodas las descargas completadas en %v\n", time.Since(inicio).Round(time.Millisecond))
}
