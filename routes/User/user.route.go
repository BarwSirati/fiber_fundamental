package user

import (
	userController "rest/api/controller/User"

	"github.com/gofiber/fiber/v2"
)

func UserInit(router fiber.Router) {
	router.Get("/", userController.GetUsers)
	router.Get("/:id", userController.GetUser)
	router.Post("/", userController.ValidateUser, userController.AddUser)
}
