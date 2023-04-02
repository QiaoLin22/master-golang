package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	Username string
	Role     string
}

func main() {
	app := fiber.New()

	app.Get("/post", handleGetPost)
	app.Get("post/manage", onlyAdmin(handleGetPostManage))
	app.Get("post/special", onlySpecialUser(handleGetPostSpecial))

	log.Fatal(app.Listen(":4000"))
}

func onlyAdmin(fn fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := getUserFromDB()
		if user.Role != "admin" {
			return c.SendStatus(http.StatusUnauthorized)
		}
		return fn(c)
	}
}

func onlySpecialUser(fn fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := getUserFromDB()
		if user.Role != "special" {
			return c.SendStatus(http.StatusUnauthorized)
		}
		return fn(c)
	}
}

func getUserFromDB() User {
	return User{
		Username: "admin",
		Role:     "user",
	}
}

func handleGetPost(c *fiber.Ctx) error {
	return c.JSON("some posts here")
}

func handleGetPostManage(c *fiber.Ctx) error {
	return c.JSON("admin page")
}

func handleGetPostSpecial(c *fiber.Ctx) error {
	return c.JSON("special posts")
}
