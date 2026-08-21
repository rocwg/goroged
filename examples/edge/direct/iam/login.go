package iam

import (
	"net/http"

	gedhttp "github.com/rocwg/ged/http"
	iamv1 "github.com/rocwg/grpc-contracts/gen/go/iam/v1"
)

func (a *Adapter) login(
	w http.ResponseWriter,
	r *http.Request,
) {
	_, err := gedhttp.DecodeJSON[iamv1.GetUserRequest](r)
	if err != nil {
		gedhttp.WriteError(w, err)
		return
	}
	// TODO ...
}
