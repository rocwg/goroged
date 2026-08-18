package dashboard

import (
	dictv1 "github.com/rocwg/grpc-contracts/gen/go/dictarea/v1"
)

type OverviewResponse struct {
	Areas   []*dictv1.AreaNode `json:"areas"`
	Message string             `json:"message"`
}
