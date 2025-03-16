package helpers

import "github.com/gofiber/fiber/v2"

type ResponseErrorOptions struct {
	StatusCode int         `json:"status_code"`
	Msg        string      `json:"msg"`
	Errors     interface{} `json:"errors"`
}

type ResponseSuccessOptions struct {
	StatusCode int         `json:"status_code"`
	Msg        string      `json:"msg"`
	Data       interface{} `json:"data"`
}

func ResponseError(c *fiber.Ctx, opts ...*ResponseErrorOptions) error {
	statusCode := fiber.StatusInternalServerError
	msg := "Something went wrong"
	var errors interface{} = nil

	if len(opts) > 0 && opts[0] != nil {
		if opts[0].StatusCode != 0 {
			statusCode = opts[0].StatusCode
		}

		if opts[0].Msg != "" {
			msg = opts[0].Msg
		}

		if opts[0].Errors != nil {
			errors = opts[0].Errors
		}
	}

	return c.Status(statusCode).JSON(
		fiber.Map{
			"status": "error",
			"msg":    msg,
			"errors": errors,
		})
}

func ResponseSuccess(c *fiber.Ctx, opts ...*ResponseSuccessOptions) error {
	statusCode := fiber.StatusOK
	msg := "Success"
	var data interface{} = nil

	if len(opts) > 0 && opts[0] != nil {
		if opts[0].StatusCode != 0 {
			statusCode = opts[0].StatusCode
		}

		if opts[0].Msg != "" {
			msg = opts[0].Msg
		}

		if opts[0].Data != nil {
			data = opts[0].Data
		}
	}

	return c.Status(statusCode).JSON(fiber.Map{
		"status": "success",
		"msg":    msg,
		"data":   data,
	})
}
