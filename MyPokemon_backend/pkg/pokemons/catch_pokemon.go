package pokemons

import (
	"math/rand"

	"github.com/gofiber/fiber/v2"
)

type CatchPokemonRequestBody struct {
	PokemonID int `json:"pokemon_id"`
}

type CatchPokemonResponseBody struct {
	PokemonID         int  `json:"pokemon_id"`
	ProbabilityChance int  `json:"probability_chance"`
	Status            bool `json:"status"`
}

func CatchPokemon(c *fiber.Ctx) error {

	data := CalculateProbability()
	status := false
	if data > 49 {
		status = true
	}

	return c.Status(fiber.StatusOK).JSON(CatchPokemonResponseBody{
		PokemonID:         1,
		ProbabilityChance: data,
		Status:            status,
	})
}

func CalculateProbability() int {
	return int(rand.Float64() * 100)
}
