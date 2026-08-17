package request

import "context"

type Context struct {
	RequestID string
	TraceID   string
}

type contextKey struct{}

func WithContext(
	ctx context.Context,
	requestContext Context,
) context.Context {
	return context.WithValue(
		ctx,
		contextKey{},
		requestContext,
	)
}

func FromContext(
	ctx context.Context,
) (Context, bool) {
	value, ok := ctx.Value(contextKey{}).(Context)
	return value, ok
}
