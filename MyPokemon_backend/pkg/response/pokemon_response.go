package response

import "phincon/pkg/common/models"

type DefaultResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// LIST POKEMON API
type PokemonListAPI struct {
	Code        int                    `json:"code"`
	Message     string                 `json:"message"`
	ListPokemon []models.ResultPokemon `json:"data"`
}

type DetailPokemonAPI struct {
	Name    string        `json:"name"`
	Sprites PokemonImages `json:"sprites"`
}
type PokemonImages struct {
	FrontDefault string `json:"front_default"`
	BackDefault  string `json:"back_default"`
	FrontShiny   string `json:"front_shiny"`
	BackShiny    string `json:"back_shiny"`
}

type PokemonListResponse struct {
	Code              int              `json:"code"`
	Message           string           `json:"message"`
	PokemonListResult []models.Pokemon `json:"data"`
}
