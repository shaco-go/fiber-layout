//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"githut.com/shaco-go/fiber-kit/core/server"
	"githut.com/shaco-go/fiber-kit/internal/handle"
	"githut.com/shaco-go/fiber-kit/internal/repo"
	"githut.com/shaco-go/fiber-kit/internal/router"
	"githut.com/shaco-go/fiber-kit/internal/service"
)

func newApp(
	http *server.HttpServer,
	task *server.TaskServer,
) *server.App {
	return server.NewApp(server.WithServer(
		http,
		task,
	))
}

func NewWire() (*server.App, func(), error) {
	panic(wire.Build(
		server.ProviderSet,
		handle.ProviderSet,
		service.ProviderSet,
		repo.ProviderSet,
		router.NewRouter,
		newApp,
	))
}
