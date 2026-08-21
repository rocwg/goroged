package iam

import (
	"encoding/json"
	"net/http"

	iamv1 "github.com/rocwg/grpc-contracts/gen/go/iam/v1"
)

func (a *Adapter) login(
	w http.ResponseWriter,
	r *http.Request,
) {
	var req iamv1.GetUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "参数格式错误", http.StatusBadRequest)
		return
	}
	// TODO ...
}
