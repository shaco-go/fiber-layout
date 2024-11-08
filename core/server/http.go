package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"githut.com/shaco-go/fiber-kit/api"
	"githut.com/shaco-go/fiber-kit/global"
	"githut.com/shaco-go/fiber-kit/internal/router"
	"net/http"
	"time"
)

func NewHttpServer(routers ...router.Router) *HttpServer {
	return &HttpServer{
		app: fiber.New(fiber.Config{
			ErrorHandler: fiberErrorHandler,
		}),
		routers: routers,
	}
}

type HttpServer struct {
	app     *fiber.App
	routers []router.Router
}

func (h *HttpServer) Start(ctx context.Context) error {

	return h.app.Listen(fmt.Sprintf(":%d", global.Conf.Port))
}

func (h *HttpServer) Stop(ctx context.Context) error {
	global.Logc.Info("Shutting down server...")

	if err := h.app.ShutdownWithTimeout(5 * time.Second); err != nil {
		global.Logc.Sugar().Fatalf("Server forced to shutdown: %+v", err)
	}

	global.Logc.Info("Server exiting")
	return nil
}

// FiberErrorHandler fiber的错误处理
func fiberErrorHandler(ctx *fiber.Ctx, err error) error {
	if err == nil {
		return nil
	}
	// 检查是否是fiber定义的错误
	var fiberErr *fiber.Error
	var myErr *api.Err
	if errors.As(err, &fiberErr) {
		if fiberErr.Code < http.StatusInternalServerError {
			// http code 小于 500,直接处理
			return api.FailWithMsg(ctx, fiberErr.Message)
		}
	} else if errors.As(err, &myErr) {
		// 检查是否自定义错误
		return ctx.JSON(api.Response{
			Code: myErr.Code(),
			Msg:  myErr.Error(),
			Data: []any{},
		})
	}
	// 如果 大于500 ,内部处理,并且上报
	return serverFail(ctx, fiberErr)
}

// serverFail 服务器内部错误
func serverFail(ctx *fiber.Ctx, err error) error {
	global.Logx.Errorf(ctx.Context(), err)
	return ctx.Status(http.StatusInternalServerError).JSON(api.Response{
		Code: api.FAIL_CODE,
		Msg:  api.FAIL_SERVER_MESSAGE,
		Data: []any{},
	})
}
