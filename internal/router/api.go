package router

import (
	"github.com/gofiber/fiber/v2"
	"githut.com/shaco-go/fiber-kit/internal/handle"
)

func NewApi(
	user *handle.User,
) *Api {
	return &Api{
		user: user,
	}
}

type Api struct {
	user *handle.User
}

func (a *Api) Register(app *fiber.App) {
	app.Get("hello-word", a.user.HelloWord)
}
