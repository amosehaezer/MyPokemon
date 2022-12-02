package request

type AddPokemonRequestBody struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}
