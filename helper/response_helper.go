package helper

import (
	"bobobox/model"
	"github.com/gofiber/fiber"
	"net/http"
)

func Ok (c *fiber.Ctx,data interface{})error{
	return c.Status(http.StatusOK).JSON(model.ResponseModel{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   data,
	})
}

func BadRequest (c *fiber.Ctx,err error)error{
	return c.Status(http.StatusBadRequest).JSON(model.ResponseModel{
		Code:   http.StatusBadRequest,
		Status: "BAD_REQUEST",
		Data:   err.Error(),
	})
}

func MethodNotAllowed(c *fiber.Ctx,err error) error {
	return c.Status(http.StatusMethodNotAllowed).JSON(model.ResponseModel{
		Code:   http.StatusMethodNotAllowed,
		Status: "METHOD_NOT_ALLOWED",
		Data:   err.Error(),
	})
}

func InternalServerError(c *fiber.Ctx,err error)error{
	return c.Status(http.StatusInternalServerError).JSON(model.ResponseModel{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL_SERVER_ERROR",
		Data:   err.Error(),
	})
}
