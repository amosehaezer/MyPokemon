package pokemons

import (
	"phincon/pkg/common/models"
	"phincon/pkg/request"

	"github.com/gofiber/fiber/v2"
)

func (h handler) AddPokemon(c *fiber.Ctx) error {
	// var pokemons []models.Pokemon
	// var res = response.PokemonListAPI{}

	body := request.AddPokemonRequestBody{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var pokemon models.Pokemon

	pokemon.Name = body.Name
	// pokemon.Image = body.Image

	if result := h.DB.Create(&pokemon); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&pokemon)
}
