package bootstrap

import (
	"context"
	"github.com/shaco-go/go-kit/logc"
	"github.com/shaco-go/go-kit/logx"
	"githut.com/shaco-go/fiber-kit/core/constant"
	"githut.com/shaco-go/fiber-kit/global"
	"go.uber.org/zap"
)

func NewLogc() *zap.Logger {
	conf := global.Conf
	c := &logc.Config{
		Env:     logc.ConvEnv(conf.Env),
		Channel: logc.ConvChannel(conf.Log.Channel),
		Level:   logc.ConvLevel(conf.Log.Level),
		Console: logc.Console{
			Level: logc.ConvLevel(conf.Log.Console.Level),
		},
		Lark: logc.Lark{
			Webhook: conf.Log.Lark.Webhook,
			Level:   logc.ConvLevel(conf.Log.Lark.Level),
		},
		File: logc.File{
			Filename:   conf.Log.File.Filename,
			MaxSize:    conf.Log.File.MaxSize,
			MaxAge:     conf.Log.File.MaxAge,
			MaxBackups: conf.Log.File.MaxBackups,
			LocalTime:  conf.Log.File.LocalTime,
			Compress:   conf.Log.File.Compress,
			Level:      logc.ConvLevel(conf.Log.File.Level),
		},
	}
	return logc.New(c)
}

func NewLogx(logger *zap.Logger) *logx.Logx {
	return logx.New(logger, func(ctx context.Context, logger *zap.Logger) *zap.Logger {
		rid, ok := ctx.Value(constant.RID).(string)
		if ok {
			logger.With(zap.String(constant.RID, rid))
		}
		return logger
	})
}
