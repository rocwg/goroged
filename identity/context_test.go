package identity_test

import (
	"context"
	"testing"

	"github.com/rocwg/ged/identity"
)

func TestContext(t *testing.T) {

	ctx := context.Background()

	expected := identity.Context{
		UserID:   "10001",
		TenantID: "demo",
	}

	ctx = identity.WithContext(
		ctx,
		expected,
	)

	actual, ok := identity.FromContext(ctx)

	if !ok {
		t.Fatal("identity context not found")
	}

	if actual != expected {
		t.Fatalf(
			"unexpected identity context: got %+v, want %+v",
			actual,
			expected,
		)
	}
}

func TestFromContext_NotFound(t *testing.T) {

	ctx := context.Background()

	_, ok := identity.FromContext(ctx)

	if ok {
		t.Fatal("expected identity context to be absent")
	}
}
