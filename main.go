package main

import (
	"encoding/json"
	"errors"
	"log"

	"os"
	"rest/api/configs"
	"rest/api/middleware"
	"rest/api/migration"
	"rest/api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
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
	Log()
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestCompression,
	}))
}

func main() {
	configs.ConnectDB()
	migration.RunMigration()
	routes.RouteInit(app.Group("/api"))
	errRun := app.Listen(":" + os.Getenv("PORT"))
	if errRun != nil {
		panic(errRun)
	}
	
}
