package main

import (
	"fmt"
	"time"
)

type Empleado struct {
	ID        int
	Nombre    string
	Puesto    string
	Salario   float64
	FechaIngreso time.Time
}

func (e Empleado) Antiguedad() int {
	return int(time.Since(e.FechaIngreso).Hours() / 24 / 365)
}

func (e Empleado) Mostrar() {
	fmt.Printf("ID: %d | %s | %s | $%.2f | Antigüedad: %d años\n",
		e.ID, e.Nombre, e.Puesto, e.Salario, e.Antiguedad())
}

func (e *Empleado) AumentarSalario(porcentaje float64) {
	e.Salario += e.Salario * porcentaje / 100
}

type Empresa struct {
	Nombre    string
	Empleados []Empleado
}

func (e *Empresa) Contratar(nombre, puesto string, salario float64) {
	id := len(e.Empleados) + 1
	emp := Empleado{
		ID:        id,
		Nombre:    nombre,
		Puesto:    puesto,
		Salario:   salario,
		FechaIngreso: time.Now(),
	}
	e.Empleados = append(e.Empleados, emp)
	fmt.Printf("Contratado: %s como %s\n", nombre, puesto)
}

func (e Empresa) ListarEmpleados() {
	fmt.Printf("\n=== Empleados de %s ===\n", e.Nombre)
	for _, emp := range e.Empleados {
		emp.Mostrar()
	}
}

func main() {
	empresa := Empresa{Nombre: "GoCorp"}

	empresa.Contratar("Alice", "Ingeniera", 75000)
	empresa.Contratar("Bob", "Diseñador", 65000)
	empresa.Contratar("Charlie", "Gerente", 90000)

	empresa.ListarEmpleados()

	fmt.Println("\n--- Aumento del 10% para todos ---")
	for i := range empresa.Empleados {
		empresa.Empleados[i].AumentarSalario(10)
	}

	empresa.ListarEmpleados()
}
