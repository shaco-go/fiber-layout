package server

import (
	"context"
	"fmt"
	"github.com/go-co-op/gocron/v2"
	"github.com/google/uuid"
	"githut.com/shaco-go/fiber-kit/global"
)

func NewTaskServer() *TaskServer {
	task := &TaskServer{}
	var err error
	task.Corn, err = gocron.NewScheduler()
	if err != nil {
		panic(fmt.Sprintf("new task server fail :%+v\n", err))
	}
	return task
}

type TaskServer struct {
	Corn gocron.Scheduler
}

func (t *TaskServer) Start(ctx context.Context) error {
	t.Corn.Start()
	return nil
}

func (t *TaskServer) Stop(ctx context.Context) error {
	return t.Corn.Shutdown()
}

// SingleTask 单例执行脚本,如果脚本还在执行,会跳过,到下一次执行的时间节点
// @cron cron的语法支持秒级,所以是6位数
// @fn 执行的脚本
// @jobName 脚本名称
func (t *TaskServer) SingleTask(cron string, fn any, jobName ...string) {
	options := []gocron.JobOption{
		gocron.WithSingletonMode(gocron.LimitModeReschedule), gocron.WithEventListeners(
			gocron.AfterJobRunsWithError(func(jobID uuid.UUID, jobName string, err error) {
				global.Logc.Sugar().Errorf("%t 执行失败\nJobID:%v\n+%+v",
					jobName,
					jobID,
					err,
				)
			}),
		),
	}
	if len(jobName) > 0 {
		options = append(options, gocron.WithName(jobName[0]))
	}
	_, err := t.Corn.NewJob(
		gocron.CronJob(cron, true),
		gocron.NewTask(fn),
		options...,
	)
	if err != nil {
		panic(err)
	}
}
