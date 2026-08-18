package request_test

import (
	"context"
	"testing"

	"github.com/rocwg/ged/request"
)

func TestContext(t *testing.T) {

	ctx := context.Background()

	expected := request.Context{
		RequestID: "req-demo-001",
		TraceID:   "trace-demo-001",
	}

	ctx = request.WithContext(
		ctx,
		expected,
	)

	actual, ok := request.FromContext(ctx)

	if !ok {
		t.Fatal("request context not found")
	}

	if actual != expected {
		t.Fatalf(
			"unexpected request context: got %+v, want %+v",
			actual,
			expected,
		)
	}
}

func TestFromContext_NotFound(t *testing.T) {

	ctx := context.Background()

	_, ok := request.FromContext(ctx)

	if ok {
		t.Fatal("expected request context to be absent")
	}
}
