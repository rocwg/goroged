package main

import (
	"context"
	"fmt"
	"time"
)

func experimentDeadline() {
	fmt.Println("Experiment 02: Context Deadline")

	ctx, cancel := context.WithTimeout(
		context.Background(),
		1*time.Second,
	)
	defer cancel()

	select {
	case <-time.After(2 * time.Second):
		fmt.Println("work finished")

	case <-ctx.Done():
		fmt.Println("context done")
		fmt.Println("error:", ctx.Err())
	}
}
