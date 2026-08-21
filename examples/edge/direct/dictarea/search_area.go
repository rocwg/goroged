package dictarea

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	dictv1 "github.com/rocwg/grpc-contracts/gen/go/dictarea/v1"
)

func (a *Adapter) searchArea(
	w http.ResponseWriter,
	r *http.Request,
) {
	var req SearchAreaRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "参数格式错误", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(
		r.Context(),
		2*time.Second,
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}
