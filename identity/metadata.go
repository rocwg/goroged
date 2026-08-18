package identity

import (
	"context"

	"google.golang.org/grpc/metadata"
)

const (
	MetadataUserID   = "x-user-id"
	MetadataTenantID = "x-tenant-id"
)

// AppendMetadata 将 Identity 写入 gRPC outgoing metadata。
//
// Identity:
//
//	HTTP Request
//	   ↓
//	context.Context
//	   ↓
//	gRPC Metadata
func AppendMetadata(
	ctx context.Context,
	identity Context,
) context.Context {

	pairs := make([]string, 0, 4)

	if identity.UserID != "" {
		pairs = append(
			pairs,
			MetadataUserID,
			identity.UserID,
		)
	}

	if identity.TenantID != "" {
		pairs = append(
			pairs,
			MetadataTenantID,
			identity.TenantID,
		)
	}

	if len(pairs) == 0 {
		return ctx
	}

	return metadata.AppendToOutgoingContext(
		ctx,
		pairs...,
	)
}
