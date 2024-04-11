package main

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Age    int    `json:"age"`
	Name   string `json:"name"`
	UserID string `json:"userid"`
}

func main() {
	app := fiber.New()

	app.Get("/user/:userid", func(c *fiber.Ctx) error {
		user := new(User)
		err := c.BodyParser(user)
		if err != nil {
			return err
		}
		fmt.Println(user.Age)
		fmt.Println(user.Name)
		fmt.Println(user.UserID)

		user.Age += 1
		userJSON, _ := json.Marshal(user)

		return c.JSON(userJSON)
	})

	app.Listen(":8888")
}
