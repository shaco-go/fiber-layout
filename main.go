package main

import (
	"context"
	"flag"
	"githut.com/shaco-go/fiber-kit/core/bootstrap"
	"githut.com/shaco-go/fiber-kit/global"
	"githut.com/shaco-go/fiber-kit/wire"
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
	// global.Db = bootstrap.NewGorm(global.Conf.Database["default"])
	// global.Redis = bootstrap.NewRedis(global.Conf.Redis)

	// 启动http
	app, f, err := wire.NewWire()
	if err != nil {
		panic(err)
	}
	defer f()
	err = app.Run(context.Background())
	if err != nil {
		panic(err)
	}
}
