package api

import "github.com/gofiber/fiber/v2"

const SUCCESS_CODE = 0
const SUCCESS_MESSAGE = "请求成功"

const FAIL_CODE = 10000
const FAIL_MESSAGE = "请求异常"
const FAIL_SERVER_MESSAGE = "服务器繁忙,请稍后再试"

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Ok(ctx *fiber.Ctx) error {
	return ctx.JSON(Response{
		Code: SUCCESS_CODE,
		Msg:  SUCCESS_MESSAGE,
		Data: []any{},
	})
}

func OkWithMsg(ctx *fiber.Ctx, message string) error {
	return ctx.JSON(Response{
		Code: SUCCESS_CODE,
		Msg:  message,
		Data: []any{},
	})
}

func OkWithData(ctx *fiber.Ctx, data any) error {
	return ctx.JSON(Response{
		Code: SUCCESS_CODE,
		Msg:  SUCCESS_MESSAGE,
		Data: data,
	})
}

func OkWithDetail(ctx *fiber.Ctx, message string, data any) error {
	return ctx.JSON(Response{
		Code: SUCCESS_CODE,
		Msg:  message,
		Data: data,
	})
}

func Fail(ctx *fiber.Ctx) error {
	return ctx.JSON(Response{
		Code: FAIL_CODE,
		Msg:  FAIL_MESSAGE,
		Data: []any{},
	})
}

func FailWithMsg(ctx *fiber.Ctx, message string) error {
	return ctx.JSON(Response{
		Code: FAIL_CODE,
		Msg:  message,
		Data: []any{},
	})
}

func FailWithData(ctx *fiber.Ctx, data any) error {
	return ctx.JSON(Response{
		Code: FAIL_CODE,
		Msg:  FAIL_MESSAGE,
		Data: data,
	})
}

func FailWithDetail(ctx *fiber.Ctx, message string, data any) error {
	return ctx.JSON(Response{
		Code: FAIL_CODE,
		Msg:  message,
		Data: data,
	})
}
