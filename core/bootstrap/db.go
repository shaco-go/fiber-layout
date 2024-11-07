package bootstrap

import (
	"context"
	"errors"
	"fmt"
	"githut.com/shaco-9696/fiber-kit/config"
	"githut.com/shaco-9696/fiber-kit/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"time"
)

func NewGorm(dsn config.DSN) *gorm.DB {
	var (
		db  *gorm.DB
		err error
	)

	db, err = gorm.Open(mysql.Open(string(dsn)), &gorm.Config{
		Logger: DefaultGormLogger(),
	})

	if err != nil {
		panic(err)
	}

	if global.IsDev {
		db = db.Debug()
	}

	// Connection Pool config
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db
}

var (
	traceStr     = "%s\n[%.3fms] [rows:%v] %s"
	traceWarnStr = "%s %s\n[%.3fms] [rows:%v] %s"
	traceErrStr  = "%s %s\n[%.3fms] [rows:%v] %s"
)

func DefaultGormLogger() *GormLogger {
	return &GormLogger{
		LogLevel:      logger.Warn,
		SlowThreshold: 200 * time.Millisecond,
	}
}

type GormLogger struct {
	LogLevel      logger.LogLevel
	SlowThreshold time.Duration
}

func (g *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	g.LogLevel = level
	return g
}

func (g *GormLogger) Info(ctx context.Context, s string, i ...interface{}) {
	if g.LogLevel >= logger.Info {
		global.Logx.WithContext(ctx).Sugar().Infof("%s\n"+s, append([]interface{}{utils.FileWithLineNum()}, i...)...)
	}
}

func (g *GormLogger) Warn(ctx context.Context, s string, i ...interface{}) {
	if g.LogLevel >= logger.Warn {
		global.Logx.WithContext(ctx).Sugar().Warnf("%s\n"+s, append([]interface{}{utils.FileWithLineNum()}, i...)...)
	}
}

func (g *GormLogger) Error(ctx context.Context, s string, i ...interface{}) {
	if g.LogLevel >= logger.Error {
		global.Logx.WithContext(ctx).Sugar().Errorf("%s\n"+s, append([]interface{}{utils.FileWithLineNum()}, i...)...)
	}
}

func (g *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if g.LogLevel <= logger.Silent {
		return
	}

	elapsed := time.Since(begin)
	switch {
	case err != nil && g.LogLevel >= logger.Error && (!errors.Is(err, logger.ErrRecordNotFound)):
		sql, rows := fc()
		if rows == -1 {
			global.Logx.WithContext(ctx).Sugar().Errorf(traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			global.Logx.WithContext(ctx).Sugar().Errorf(traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case elapsed > g.SlowThreshold && g.SlowThreshold != 0 && g.LogLevel >= logger.Warn:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", g.SlowThreshold)
		if rows == -1 {
			global.Logx.WithContext(ctx).Sugar().Warnf(traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			global.Logx.WithContext(ctx).Sugar().Warnf(traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case g.LogLevel == logger.Info:
		sql, rows := fc()
		if rows == -1 {
			global.Logx.WithContext(ctx).Sugar().Infof(traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			global.Logx.WithContext(ctx).Sugar().Infof(traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	}
}
