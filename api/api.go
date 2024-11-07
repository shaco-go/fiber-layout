package api

import "github.com/gofiber/fiber/v2"

const SUCCESS_CODE = 0
const SUCCESS_MESSAGE = "SUCCESS"

const FAIL_CODE = 10000
const FAIL_MESSAGE = "FAIL"

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
