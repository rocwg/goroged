package grpc_test

import (
	"context"
	"testing"

	"github.com/rocwg/ged/grpc"
	"github.com/rocwg/ged/identity"
	"github.com/rocwg/ged/request"
	"google.golang.org/grpc/metadata"
)

func TestPropagateMetadata(t *testing.T) {

	ctx := context.Background()

	ctx = identity.WithContext(
		ctx,
		identity.Context{
			UserID:   "10001",
			TenantID: "demo",
		},
	)

	ctx = request.WithContext(
		ctx,
		request.Context{
			RequestID: "req-demo-001",
			TraceID:   "trace-demo-001",
		},
	)

	ctx = grpc.PropagateMetadata(ctx)

	md, ok := metadata.FromOutgoingContext(ctx)

	if !ok {
		t.Fatal("outgoing metadata not found")
	}

	assertMetadata(t, md, "x-user-id", "10001")
	assertMetadata(t, md, "x-tenant-id", "demo")
	assertMetadata(t, md, "x-request-id", "req-demo-001")
	assertMetadata(t, md, "x-trace-id", "trace-demo-001")
}

func TestPropagateMetadata_EmptyContext(t *testing.T) {

	ctx := context.Background()

	ctx = grpc.PropagateMetadata(ctx)

	md, ok := metadata.FromOutgoingContext(ctx)

	if ok && len(md) != 0 {
		t.Fatalf(
			"expected empty metadata, got %v",
			md,
		)
	}
}

func assertMetadata(
	t *testing.T,
	md metadata.MD,
	key string,
	expected string,
) {

	t.Helper()

	values := md.Get(key)

	if len(values) != 1 {
		t.Fatalf(
			"metadata %q: got %v, want one value",
			key,
			values,
		)
	}

	if values[0] != expected {
		t.Fatalf(
			"metadata %q: got %q, want %q",
			key,
			values[0],
			expected,
		)
	}
}
