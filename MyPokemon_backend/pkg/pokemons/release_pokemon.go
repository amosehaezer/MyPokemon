package pokemons

import (
	"math"
	"math/rand"

	"github.com/gofiber/fiber/v2"
)

type ReleasePokemonResponse struct {
	PokemonName    string `json:"pokemon_name"`
	ResponseNumber int    `json:"response_number"`
	ReleaseStatus  bool   `json:"release_status"`
}

func ReleasePokemon(c *fiber.Ctx) error {
	// id := c.Params("id")

	num := RandomReleaseNumber()
	return c.Status(fiber.StatusOK).JSON(ReleasePokemonResponse{
		PokemonName:    "",
		ResponseNumber: num,
		ReleaseStatus:  CheckPrimeNumber(num),
	})
}

func RandomReleaseNumber() int {
	return rand.Intn(100)
}

func CheckPrimeNumber(num int) bool {
	if num < 2 {
		return false
	}
	sq_root := int(math.Sqrt(float64(num)))
	for i := 2; i <= sq_root; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
