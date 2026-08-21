package hello

import (
	"context"
	"encoding/json"
	"net/http"

	hellov1 "github.com/rocwg/grpc-contracts/gen/go/hello/v1"
)

func (a *Adapter) sayHello(
	w http.ResponseWriter,
	r *http.Request,
) {
	var req SayHelloRequest

	req.Name = r.URL.Query().Get("name")

	if req.Name == "" {
		req.Name = "World"
	}

	ctx, cancel := context.WithTimeout(
		r.Context(),
		a.unaryTimeout,
	)
	defer cancel()

	resp, err := a.client.SayHello(
		ctx,
		&hellov1.SayHelloRequest{
			Name: req.Name,
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
