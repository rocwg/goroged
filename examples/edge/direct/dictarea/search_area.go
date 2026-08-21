package dictarea

import (
	"context"
	"net/http"

	gedhttp "github.com/rocwg/ged/http"
	dictv1 "github.com/rocwg/grpc-contracts/gen/go/dictarea/v1"
)

func (a *Adapter) searchArea(
	w http.ResponseWriter,
	r *http.Request,
) {
	req, err := gedhttp.DecodeJSON[SearchAreaRequest](r)
	if err != nil {
		gedhttp.WriteError(w, err)
		return
	}

	ctx, cancel := context.WithTimeout(
		r.Context(),
		a.unaryTimeout,
	)
	defer cancel()

	resp, err := a.client.SearchArea(
		ctx,
		&dictv1.SearchAreaRequest{
			Keyword: req.Keyword,
			Limit:   req.Limit,
		},
	)

	if err != nil {
		gedhttp.WriteError(w, err)
		return
	}
	_ = gedhttp.WriteJSON(w, http.StatusOK, resp)
}
