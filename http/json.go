package http

import (
	"encoding/json"
	nethttp "net/http"
)

// DecodeJSON 将 HTTP JSON Request Body 解码为 T。
//
// T 是当前 HTTP 请求对应的 Request Model。
func DecodeJSON[T any](
	r *nethttp.Request,
) (T, error) {

	var req T

	err := json.NewDecoder(r.Body).Decode(&req)

	return req, err
}
