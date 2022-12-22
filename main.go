package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type User struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type IError struct {
	Field string
	Tag   string
	Value string
}

var validate = validator.New()

func ValidateUser(c *fiber.Ctx) error {
	var errors []*IError
	body := new(User)
	c.BodyParser(&body)

	err := validate.Struct(body)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var el IError
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			errors = append(errors, &el)
		}
		return c.SendStatus(fiber.ErrBadRequest.Code)
	}
	return c.Next()
}

func main() {
	app := fiber.New(fiber.Config{
		Prefork:      true,
		ServerHeader: "Fiber",
		AppName:      "API v0.1",
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	file, err := os.OpenFile("./log/tmp.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()
	app.Use(logger.New(logger.Config{
		TimeZone: "Asia/Bangkok",
		Format:   "[${ip}]:${port} ${status} - ${method} ${path}\n",
		Output:   file,
	}))

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestCompression,
	}))

	api := app.Group("/api")
	user := api.Group("/user")

	user.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	user.Post("/", ValidateUser, func(c *fiber.Ctx) error {
		fmt.Println(string(c.Request().Body()))
		data := User{
			Username: "bxdman",
			Password: "bxdman",
		}
		return c.JSON(data)
	})
	log.Fatal((app.Listen(":3000")))
}
