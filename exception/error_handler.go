package exception

import (
	"bobobox/helper"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	status, _ := err.(ValidationError)

	if status.Status == http.StatusBadRequest {
		return helper.BadRequest(ctx,err)
	}

	if e, ok := err.(*fiber.Error); ok {
		if e.Code == http.StatusMethodNotAllowed {
			return helper.MethodNotAllowed(ctx,err)
		}
	}

	return helper.InternalServerError(ctx,err)
}
