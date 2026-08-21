package main

import "fmt"

func StringFirst(values []string) string {
	if len(values) == 0 {
		return ""
	}

	return values[0]
}

func IntFirst(values []int) int {
	if len(values) == 0 {
		return 0
	}

	return values[0]
}

func experiment04() {

	names := []string{"roc", "goro", "ged"}
	numbers := []int{10, 20, 30}

	fmt.Println("Experiment 04")
	fmt.Println(StringFirst(names))
	fmt.Println(IntFirst(numbers))
}
