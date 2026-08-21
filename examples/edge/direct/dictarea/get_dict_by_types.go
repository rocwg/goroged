package dictarea

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	dictv1 "github.com/rocwg/grpc-contracts/gen/go/dictarea/v1"
)

func (a *Adapter) getDictByTypes(
	w http.ResponseWriter,
	r *http.Request,
) {
	var req GetDictByTypesRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "参数格式错误", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(
		r.Context(),
		2*time.Second,
	)
	defer cancel()

	resp, err := a.client.GetDictByTypes(
		ctx,
		&dictv1.GetDictByTypesRequest{
			TypeCodes: req.TypeCodes,
		},
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}
