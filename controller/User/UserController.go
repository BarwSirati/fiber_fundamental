package user

import (
	"log"
	"net/http"
	"rest/api/configs"
	ResType "rest/api/models/Response"
	UserModel "rest/api/models/User"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type (
	User     UserModel.User
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

func GetUsers(c *fiber.Ctx) error {
	var users []UserModel.User
	result := configs.DB.Select("id", "username", "name", "lastname").Find(&users)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return c.JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	var users []UserModel.User
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	result := configs.DB.First(&users, "id = ?", id)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return c.JSON(users)
}

func AddUser(c *fiber.Ctx) error {
	user := new(UserModel.User)
	if err := c.BodyParser(user); err != nil {
		return err
	}
	newUser := UserModel.User{
		Name:     user.Name,
		Lastname: user.Lastname,
		Username: user.Username,
		Password: user.Password,
	}

	errCreateUser := configs.DB.Create(&newUser).Error

	if errCreateUser != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "500 Internal Server Error",
		})
	}
	res := Response{
		Data:   http.StatusText(http.StatusCreated),
		Status: fiber.StatusCreated,
	}
	return c.JSON(res)
}
