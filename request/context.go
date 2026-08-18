package request

import "context"

// Context 表示当前 Edge Request 的上下文信息。
type Context struct {
	RequestID string
	TraceID   string
}

type contextKey struct{}

// WithContext 将 Request Context 写入 context。
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

// FromContext 从 context 中读取 Request Context。
func FromContext(
	ctx context.Context,
) (Context, bool) {
	requestContext, ok := ctx.Value(contextKey{}).(Context)
	return requestContext, ok
}
