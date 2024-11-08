package server

import (
	"context"
	"github.com/google/wire"
)

type Server interface {
	// Start ...
	Start(ctx context.Context) error
	// Stop ...
	Stop(ctx context.Context) error
}

var ProviderSet = wire.NewSet(NewHttpServer, NewTaskServer)
