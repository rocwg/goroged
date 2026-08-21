package main

import "fmt"

type Result[T any] struct {
	Value T
	Err   error
}

func Success[T any](value T) Result[T] {
	return Result[T]{
		Value: value,
	}
}

func Failure[T any](err error) Result[T] {
	return Result[T]{
		Err: err,
	}
}

func experiment02() {

	stringResult := Success("hello")
	intResult := Success(42)

	fmt.Println("Experiment 02")

	fmt.Printf(
		"stringResult: value=%q error=%v\n",
		stringResult.Value,
		stringResult.Err,
	)

	fmt.Printf(
		"intResult: value=%d error=%v\n",
		intResult.Value,
		intResult.Err,
	)
}
