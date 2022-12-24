package routes

import (
	userRoute "rest/api/routes/User"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(api fiber.Router) {
	userRoute.UserInit(api.Group("/user"))
}
