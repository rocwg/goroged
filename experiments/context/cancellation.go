package main

import (
	"context"
	"fmt"
	"time"
)

func experimentCancellation() {
	fmt.Println("Experiment 01: Context Cancellation")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	done := make(chan struct{})

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("worker: cancelled")
				close(done)
				return

			default:
				fmt.Println("worker: working...")
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()

	time.Sleep(1 * time.Second)

	fmt.Println("main: cancel context")
	cancel()

	<-done

	fmt.Println("main: worker stopped")
}
