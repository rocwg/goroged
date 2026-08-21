package dictarea

import (
	"context"
	"net/http"

	gedhttp "github.com/rocwg/ged/http"
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
		a.unaryTimeout,
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
		gedhttp.WriteError(w, err)
		return
	}
	_ = gedhttp.WriteJSON(w, http.StatusOK, resp)
}
