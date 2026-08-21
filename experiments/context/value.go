package main

import (
	"context"
	"fmt"
)

type contextKey string

const requestIDKey contextKey = "request_id"

func experimentValue() {
	fmt.Println("Experiment 03: Context Value")

	ctx := context.Background()

	ctx = context.WithValue(
		ctx,
		requestIDKey,
		"req-demo-001",
	)

	requestID, ok := ctx.Value(requestIDKey).(string)
	if !ok {
		fmt.Println("request_id not found")
		return
	}

	fmt.Println("request_id:", requestID)
}
