package bootstrap

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"githut.com/shaco-go/fiber-kit/global"
)

func Viper(in string) *viper.Viper {
	v := viper.New()
	v.SetConfigFile(in)
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = v.Unmarshal(&global.Conf)
	if err != nil {
		panic(err)
	}
	// 文件热更新
	v.OnConfigChange(func(in fsnotify.Event) {
		err = v.Unmarshal(&global.Conf)
		if err != nil {
			panic(err)
		}
	})
	v.WatchConfig()
	return v
}
