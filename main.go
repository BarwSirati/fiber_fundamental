package main

import (
	"encoding/json"
	"errors"
	"log"

	"os"
	"rest/api/middleware"
	"rest/api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var (
	app = fiber.New(fiber.Config{
		Prefork:      true,
		ServerHeader: "Fiber",
		AppName:      "API v0.1",
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
	})
)

func init() {

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestCompression,
	}))
}

func Log() {
	path := "log"
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}

	file, err := os.OpenFile("./log/tmp.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	app.Use(logger.New(logger.Config{
		TimeZone: "Asia/Bangkok",
		Format:   "[${ip}]:${port} ${status} - ${method} ${path}\n",
		Output:   file,
	}))

	app.Use(middleware.New())
}

func main() {
	Log()
	api := app.Group("/api")
	routes.RouteInit(api)
	errRun := app.Listen(":3000")
	if errRun != nil {
		panic(errRun)
	}
}
