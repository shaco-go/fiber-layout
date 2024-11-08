package handle

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"githut.com/shaco-go/fiber-kit/internal/service"
	"time"
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
	result := ctx.Params("aa")
	fmt.Println(result)
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println(result)
	}()
	return ctx.SendString(u.us.GetUser())
}

func (u *User) Test1(ctx *fiber.Ctx) error {
	result := ctx.Params("aa")
	fmt.Println(result)
	return ctx.SendString("test")
}
