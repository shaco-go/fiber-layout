package global

import (
	"github.com/redis/go-redis/v9"
	"github.com/shaco-go/go-kit/logx"
	"github.com/spf13/viper"
	"githut.com/shaco-go/fiber-kit/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Conf  *config.Config
	Viper *viper.Viper
	Logc  *zap.Logger
	Logx  *logx.Logx
	Db    *gorm.DB
	IsDev bool
	Redis *redis.Client
)
