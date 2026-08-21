package main

import "fmt"

// Map 将 []T 转换成 []R。
//
// T：输入元素类型
// R：输出元素类型
func Map[T any, R any](
	values []T,
	fn func(T) R,
) []R {

	result := make([]R, 0, len(values))

	for _, value := range values {
		result = append(result, fn(value))
	}

	return result
}

func experiment01() {

	names := []string{
		"roc",
		"goro",
		"ged",
	}

	lengths := Map(
		names,
		func(name string) int {
			return len(name)
		},
	)

	fmt.Println("Experiment 01")
	fmt.Println(names)
	fmt.Println(lengths)
}
