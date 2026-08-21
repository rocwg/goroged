package dictarea

import (
	"context"
	"net/http"

	gedhttp "github.com/rocwg/ged/http"
	dictv1 "github.com/rocwg/grpc-contracts/gen/go/dictarea/v1"
)

func (a *Adapter) batchGetAreaByCodes(
	w http.ResponseWriter,
	r *http.Request,
) {
	req, err := gedhttp.DecodeJSON[BatchGetAreaByCodesRequest](r)
	if err != nil {
		gedhttp.WriteError(w, err)
		return
	}

	ctx, cancel := context.WithTimeout(
		r.Context(),
		a.unaryTimeout,
	)
	defer cancel()

	resp, err := a.client.BatchGetAreaByCodes(
		ctx,
		&dictv1.BatchGetAreaByCodesRequest{
			AreaCodes: req.AreaCodes,
		},
	)

	if err != nil {
		gedhttp.WriteError(w, err)
		return
	}
	_ = gedhttp.WriteJSON(w, http.StatusOK, resp)
}
