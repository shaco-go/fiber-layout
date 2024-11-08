package handle

import (
	"github.com/gofiber/fiber/v2"
	"githut.com/shaco-go/fiber-kit/internal/service"
)

func NewUser(
	us *service.User,
) *User {
	return &User{
		us: us,
	}
}

type User struct {
	us *service.User
}

func (u *User) HelloWord(ctx *fiber.Ctx) error {
	return ctx.SendString(u.us.GetUser())
}
