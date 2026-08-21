package hello

import (
	"context"
	"net/http"

	gedhttp "github.com/rocwg/ged/http"
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
		gedhttp.WriteError(w, err)
		return
	}
	_ = gedhttp.WriteJSON(w, http.StatusOK, resp)
}
