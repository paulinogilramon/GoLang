package main

import "fmt"

// Stack[T] - tipo genérico
type Stack[T any] struct {
	elementos []T
}

func (s *Stack[T]) Push(elem T) {
	s.elementos = append(s.elementos, elem)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.elementos) == 0 {
		var zero T
		return zero, false
	}
	elem := s.elementos[len(s.elementos)-1]
	s.elementos = s.elementos[:len(s.elementos)-1]
	return elem, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if len(s.elementos) == 0 {
		var zero T
		return zero, false
	}
	return s.elementos[len(s.elementos)-1], true
}

func (s *Stack[T]) Size() int {
	return len(s.elementos)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.elementos) == 0
}

// Constraint personalizada
type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

func Sum[T Numeric](stack *Stack[T]) T {
	var total T
	for _, v := range stack.elementos {
		total += v
	}
	return total
}

func main() {
	// Stack de enteros
	intStack := Stack[int]{}
	intStack.Push(10)
	intStack.Push(20)
	intStack.Push(30)

	fmt.Println("Stack[int]:")
	for !intStack.IsEmpty() {
		if val, ok := intStack.Pop(); ok {
			fmt.Printf("  Pop: %d\n", val)
		}
	}

	// Stack de strings
	strStack := Stack[string]{}
	strStack.Push("Go")
	strStack.Push("Rust")
	strStack.Push("Java")

	fmt.Println("\nStack[string]:")
	if val, ok := strStack.Peek(); ok {
		fmt.Println("  Peek:", val)
	}
	for !strStack.IsEmpty() {
		if val, ok := strStack.Pop(); ok {
			fmt.Printf("  Pop: %s\n", val)
		}
	}

	// Sum con genéricos y constraint Numeric
	numStack := Stack[int]{}
	numStack.Push(5)
	numStack.Push(10)
	numStack.Push(15)
	fmt.Println("\nSuma del stack:", Sum(&numStack))

	// Pop de stack vacío
	empty := Stack[float64]{}
	if _, ok := empty.Pop(); !ok {
		fmt.Println("Stack vacío: Pop devolvió false")
	}
}
