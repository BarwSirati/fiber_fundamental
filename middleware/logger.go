package middleware

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
)

func New() fiber.Handler {
	return func(c *fiber.Ctx) error {
		before := time.Now()
		if err := c.Next(); err != nil {
			return err
		}
		diff := time.Since(before)
		log.Printf("%d | %s | %s | %v", c.Response().StatusCode(),
			c.Method(), c.Path(), diff)
		return nil
	}
}
