package request

import (
	"context"

	"google.golang.org/grpc/metadata"
)

const (
	MetadataRequestID = "x-request-id"
	MetadataTraceID   = "x-trace-id"
)

func AppendMetadata(
	ctx context.Context,
	requestContext Context,
) context.Context {

	pairs := make([]string, 0, 4)

	if requestContext.RequestID != "" {
		pairs = append(
			pairs,
			MetadataRequestID,
			requestContext.RequestID,
		)
	}

	if requestContext.TraceID != "" {
		pairs = append(
			pairs,
			MetadataTraceID,
			requestContext.TraceID,
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
