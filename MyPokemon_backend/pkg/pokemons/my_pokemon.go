package pokemons

import (
	"phincon/pkg/response"

	"github.com/gofiber/fiber/v2"
)

func ListMyPokemon(c *fiber.Ctx) error {
	response := response.DefaultResponse{
		Code:    200,
		Message: "Successfully get my Pokemon List",
	}

	return c.Status(fiber.StatusOK).JSON(&response)
}
