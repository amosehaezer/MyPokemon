package pokemons

import (
	"phincon/pkg/response"

	"github.com/gofiber/fiber/v2"
)

func RenamePokemon(c *fiber.Ctx) error {

	response := response.DefaultResponse{
		Code:    200,
		Message: "Successfully get my Pokemon List",
	}

	return c.Status(fiber.StatusOK).JSON(&response)
}

func (h handler) GenerateFibonacci() string {
	return ""
}

func FibonacciRecursion(n int) int {
	if n <= 1 {
		return n
	}
	return FibonacciRecursion(n-1) + FibonacciRecursion(n-2)
}
