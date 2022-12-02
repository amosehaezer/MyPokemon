package pokemons

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	h := &handler{
		DB: db,
	}
	routes := app.Group("/api")

	// LIST POKEMON
	routes.Get("/", h.GetPokemons) // DATABASE
	routes.Get("/pokemon", h.GetListPokemons)
	routes.Get("/pokemon/:id", h.GetAPIPokemonDetail) // API
	routes.Get("/pokemon/seed", h.GetAPIListPokemons) // API

	routes.Post("/", h.AddPokemon) // DATABASE

	routes.Get("/catch/:id", CatchPokemon)
	routes.Get("/release/:id", ReleasePokemon)
	routes.Get("/rename/:id", RenamePokemon)

	routes.Get("/pokemon/my", ListMyPokemon)
}
