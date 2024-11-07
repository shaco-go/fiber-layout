package main

import (
	"flag"
	"githut.com/shaco-9696/fiber-kit/core/bootstrap"
	"githut.com/shaco-9696/fiber-kit/global"
	"strings"
)

func main() {
	var configFile = flag.String("conf", "dev.yaml", "配置项")
	flag.Parse()

	// 初始化项目
	global.Viper = bootstrap.Viper(*configFile)
	global.IsDev = strings.ToLower(global.Conf.Env) == "dev"
	global.Logc = bootstrap.NewLogc()
	global.Logx = bootstrap.NewLogx(global.Logc)
	global.Db = bootstrap.NewGorm(global.Conf.Database["default"])
	global.Redis = bootstrap.NewRedis(global.Conf.Redis)

	// 启动http

	select {}
	// app := fiber.New(fiber.Config{
	// 	ErrorHandler: func(ctx *fiber.Ctx, err error) error {
	// 		fmt.Printf("err:%+v", err)
	// 		return nil
	// 	},
	// })
	// fiber.NewError()
	// app.Get("/test", func(ctx *fiber.Ctx) error {
	// 	_, err := ctx.Write([]byte("test"))
	// 	return err
	// })
	// app.Get("/test1", func(ctx *fiber.Ctx) error {
	// 	return xerror.New("ceshi")
	// })
	// app.Get("/test2", func(ctx *fiber.Ctx) error {
	// 	panic("ceshi22222")
	// 	return nil
	// })
	// panic(app.Listen(":8989"))
}
