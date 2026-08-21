package dictarea

import (
	"net/http"

	dictv1 "github.com/rocwg/grpc-contracts/gen/go/dictarea/v1"
)

func Register(
	mux *http.ServeMux,
	client dictv1.DictAreaServiceClient,
) {
	adapter := NewAdapter(client)

	mux.HandleFunc("GET /api/v1/dictarea/get-by-parent", adapter.getAreaByParent)
	mux.HandleFunc("POST /api/v1/dictarea/search", adapter.searchArea)
	mux.HandleFunc("POST /api/v1/dictarea/batch-get-by-codes", adapter.batchGetAreaByCodes)
	mux.HandleFunc("POST /api/v1/dictarea/dicts", adapter.getDictByTypes)
}
