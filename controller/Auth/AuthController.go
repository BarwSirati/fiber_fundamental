package auth

import (
	"net/http"
	login "rest/api/models/Login"
	ResType "rest/api/models/Response"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type (
	Login    login.Login
	Response ResType.Response
)

var validate = validator.New()

func ValidateAuth(c *fiber.Ctx) error {
	body := new(Login)
	c.BodyParser(&body)

	err := validate.Struct(body)
	if err != nil {
		res := Response{
			Data:   http.StatusText(http.StatusBadRequest),
			Status: fiber.ErrBadRequest.Code,
		}
		return c.JSON(res)
	}
	return c.Next()
}
