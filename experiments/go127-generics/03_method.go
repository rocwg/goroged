package main

import "fmt"

type Box[T any] struct {
	value T
}

func NewBox[T any](value T) Box[T] {
	return Box[T]{
		value: value,
	}
}

func (b Box[T]) Value() T {
	return b.value
}

func experiment03() {

	stringBox := NewBox("hello")
	intBox := NewBox(42)

	fmt.Println("Experiment 03")
	fmt.Printf("stringBox: %s\n", stringBox.Value())
	fmt.Printf("intBox: %d\n", intBox.Value())
}
