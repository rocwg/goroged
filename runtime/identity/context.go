package identity

import "context"

type Context struct {
	UserID   string
	TenantID string
}

type contextKey struct{}

func WithContext(ctx context.Context, identity Context) context.Context {
	return context.WithValue(ctx, contextKey{}, identity)
}

func FromContext(ctx context.Context) (Context, bool) {
	identity, ok := ctx.Value(contextKey{}).(Context)
	return identity, ok
}
