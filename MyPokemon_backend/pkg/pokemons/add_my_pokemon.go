package pokemons

import (
	"phincon/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type AddPokemonRequest struct {
	PokemonID   int    `json:"pokemon_id"`
	PokemonName string `json:"pokemon_name"`
}

func (h handler) AddMyPokemon(c *fiber.Ctx) error {
	response := response.PokemonListResponse{
		Code:    200,
		Message: "Successfully get the Pokemon List",
		// PokemonListResult: pokemons,
	}
	return c.Status(fiber.StatusOK).JSON(&response)
}
