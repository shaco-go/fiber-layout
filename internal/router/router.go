package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"githut.com/shaco-go/fiber-kit/core/constant"
	"githut.com/shaco-go/fiber-kit/internal/handle"
)

func NewRouter(user *handle.User) *Router {
	return &Router{
		user: user,
	}
}

type Router struct {
	user *handle.User
}

func (r *Router) Register(app *fiber.App) {
	a := app.Group("/", requestid.New(requestid.Config{
		ContextKey: constant.RID,
	}))
	a.Get("test1/:aa", r.user.HelloWord)
	a.Get("test2/:aa", r.user.Test1)
}
