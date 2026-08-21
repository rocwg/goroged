package dictarea

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	dictv1 "github.com/rocwg/grpc-contracts/gen/go/dictarea/v1"
)

func (a *Adapter) getAreaByParent(
	w http.ResponseWriter,
	r *http.Request,
) {
	// 1. Bind Request
	var req GetAreaByParentRequest

	req.ParentCode = r.URL.Query().Get("parentCode")

	if req.ParentCode == "" {
		req.ParentCode = "0"
	}

	// 2. context.WithTimeout()
	ctx, cancel := context.WithTimeout(
		r.Context(),
		2*time.Second,
	)
	defer cancel()

	// 3. grpc Request
	// 4. grpc Call
	resp, err := a.client.GetAreaByParent(
		ctx,
		&dictv1.GetAreaByParentRequest{
			ParentCode: req.ParentCode,
		},
	)

	// 5. Return JSON
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}
