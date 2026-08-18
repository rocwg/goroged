package authentication

import gedidentity "github.com/rocwg/ged/identity"

// Authenticator 将认证凭证解析并转换为 Edge Identity。
//
// ged 只定义认证能力的接入点，
// 不实现具体认证机制。
type Authenticator interface {
	Authenticate(token string) (gedidentity.Context, error)
}
