package user

import (
	"net/http"
	ResType "rest/api/interfaces/Response"
	UserType "rest/api/interfaces/User"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type (
	User     UserType.User
	Response ResType.Response
)

var validate = validator.New()

func ValidateUser(c *fiber.Ctx) error {
	body := new(User)
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

func GetUser(c *fiber.Ctx) error {
	data := User{
		Username: "test",
		Password: "bxdman",
	}
	return c.JSON(data)
}

func AddUser(c *fiber.Ctx) error {

	return c.Send(c.Body())
}
