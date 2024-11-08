package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

type Router interface {
	Register(app *fiber.App)
}

func NewRouter(
	api *Api,
) []Router {
	return []Router{api}
}

var ProviderSet = wire.NewSet(
	NewApi,
	NewRouter,
)
