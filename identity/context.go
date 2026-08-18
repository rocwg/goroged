package identity

import "context"

// Context 表示当前 HTTP Request 对应的身份上下文。
//
// Identity 由 Authentication 建立，
// 后续通过 context.Context 在 Edge 内部传播。
type Context struct {
	UserID   string
	TenantID string
}

type contextKey struct{}

// WithContext 将 Identity 写入 context。
func WithContext(
	ctx context.Context,
	identity Context,
) context.Context {
	return context.WithValue(
		ctx,
		contextKey{},
		identity,
	)
}

// FromContext 从 context 中读取 Identity。
func FromContext(
	ctx context.Context,
) (Context, bool) {
	identity, ok := ctx.Value(contextKey{}).(Context)
	return identity, ok
}
