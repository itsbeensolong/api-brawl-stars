package services

import (
	"encoding/json"
	"os"

	"github.com/gofiber/fiber/v2"
)

func App(c *fiber.Ctx) error {
	data, err := os.ReadFile("main.json")

	if err != nil {
		return err
	}

	var jsonData interface{}

	err = json.Unmarshal(data, &jsonData)

	if err != nil {
		return err
	}

	res, err := json.Marshal(jsonData)

	if err != nil {
		return err
	}

	return c.Send(res)
}
