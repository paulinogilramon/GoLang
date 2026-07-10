package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type Persona struct {
	Nombre string `json:"nombre"`
	Edad   int    `json:"edad"`
	Activo bool   `json:"activo"`
}

func main() {
	// --- fmt ---
	fmt.Println("=== fmt ===")
	nombre := "Alice"
	edad := 30
	fmt.Printf("Hola %s, tienes %d años\n", nombre, edad)
	s := fmt.Sprintf("Formateado: %s - %d", nombre, edad)
	fmt.Println(s)

	// --- strings ---
	fmt.Println("\n=== strings ===")
	txt := "  Hola, Mundo!  "
	fmt.Println("Trim:", strings.TrimSpace(txt))
	fmt.Println("Upper:", strings.ToUpper(txt))
	fmt.Println("Contains 'Mundo':", strings.Contains(txt, "Mundo"))
	fmt.Println("Split:", strings.Split("a,b,c", ","))
	fmt.Println("Join:", strings.Join([]string{"x", "y", "z"}, "-"))

	// --- strconv ---
	fmt.Println("\n=== strconv ===")
	num, _ := strconv.Atoi("42")
	fmt.Println("Atoi:", num)
	str := strconv.Itoa(123)
	fmt.Println("Itoa:", str)
	flt, _ := strconv.ParseFloat("3.14", 64)
	fmt.Println("ParseFloat:", flt)

	// --- time ---
	fmt.Println("\n=== time ===")
	now := time.Now()
	fmt.Println("Ahora:", now.Format(time.RFC3339))
	fmt.Println("Año:", now.Year())
	fmt.Println("Mes:", now.Month())

	later := now.Add(2 * time.Hour)
	fmt.Println("Diferencia:", later.Sub(now))

	// --- math ---
	fmt.Println("\n=== math ===")
	fmt.Println("Pi:", math.Pi)
	fmt.Println("Sqrt(16):", math.Sqrt(16))
	fmt.Println("Pow(2,10):", math.Pow(2, 10))
	fmt.Println("Round(3.7):", math.Round(3.7))
	fmt.Println("Max(10,20):", math.Max(10, 20))

	// --- encoding/json ---
	fmt.Println("\n=== encoding/json ===")
	p := Persona{Nombre: "Bob", Edad: 25, Activo: true}
	jsonBytes, _ := json.Marshal(p)
	fmt.Println("JSON:", string(jsonBytes))

	jsonStr := `{"nombre":"Charlie","edad":35,"activo":false}`
	var p2 Persona
	json.Unmarshal([]byte(jsonStr), &p2)
	fmt.Printf("Desde JSON: %+v\n", p2)

	// --- os, bufio, io ---
	fmt.Println("\n=== os / bufio / io ===")
	archivo := "ejemplo_std.txt"
	contenido := "línea 1\nlínea 2\nlínea 3\n"
	os.WriteFile(archivo, []byte(contenido), 0644)
	defer os.Remove(archivo)

	file, _ := os.Open(archivo)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println("Scanner:", scanner.Text())
	}

	file.Seek(0, io.SeekStart)
	bytes, _ := io.ReadAll(file)
	fmt.Println("ReadAll:", string(bytes))
}
