package hello

import (
	bridgegrpc "github.com/rocwg/goro-edge/bridge/internal/grpc"
)

type Adapter struct {
	clients *bridgegrpc.Clients
}

func NewAdapter(clients *bridgegrpc.Clients) *Adapter {
	return &Adapter{
		clients: clients,
	}
}
