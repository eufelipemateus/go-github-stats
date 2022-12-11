package handlers

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	err = ctx.Status(code).Render(fmt.Sprintf("errors/%d", code), fiber.Map{ "Message": err.Error()   })
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error"+err.Error())
	}
	return nil
}