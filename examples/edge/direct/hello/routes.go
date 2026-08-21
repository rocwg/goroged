package hello

import (
	"net/http"

	hellov1 "github.com/rocwg/grpc-contracts/gen/go/hello/v1"
)

func Register(
	mux *http.ServeMux,
	client hellov1.HelloServiceClient,
) {
	adapter := NewAdapter(client)

	mux.HandleFunc("GET /api/v1/hello/unary", adapter.sayHello)
	mux.HandleFunc("GET /api/v1/hello/stream", adapter.streamHello)
}
