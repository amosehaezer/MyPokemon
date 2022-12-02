package pokemons

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"phincon/pkg/common/models"
	"phincon/pkg/response"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
)

func (h handler) GetPokemons(c *fiber.Ctx) error {
	var pokemons []models.Pokemon

	if result := h.DB.Find(&pokemons); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}
	response := response.PokemonListResponse{
		Code:              200,
		Message:           "Successfully get the Pokemon List",
		PokemonListResult: pokemons,
	}

	return c.Status(fiber.StatusOK).JSON(&response)
}

// CURL API LIST POKEMON
func (h handler) GetAPIListPokemons(c *fiber.Ctx) error {

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://pokeapi.co/api/v2/pokemon", nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var jsonData = []byte(bodyText)
	var list models.ListPokemon

	err = json.Unmarshal(jsonData, &list)
	if err != nil {
		fmt.Println(err.Error())
	}

	var res []models.ResultPokemon
	copier.Copy(&res, &list.Result)

	response := response.PokemonListAPI{
		Code:        200,
		Message:     "Successfully get the Pokemon List",
		ListPokemon: res,
	}

	return c.Status(fiber.StatusOK).JSON(&response)
}

// CURL API POKEMON DETAIL
func (h handler) GetAPIPokemonDetail(c *fiber.Ctx) error {
	id := c.Params("id")

	client := &http.Client{}
	link := fmt.Sprintf("%s/%s", "https://pokeapi.co/api/v2/pokemon", id)
	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyText, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s\n", bodyText)
	if err != nil {
		log.Fatal(err)
	}

	var jsonData = []byte(bodyText)
	var detail response.DetailPokemonAPI

	err = json.Unmarshal(jsonData, &detail)
	if err != nil {
		fmt.Println(err.Error())
	}

	var res response.PokemonImages
	copier.Copy(&res, &detail.Sprites)

	response := response.DetailPokemonAPI{
		Name:    detail.Name,
		Sprites: res,
	}
	return c.Status(fiber.StatusOK).JSON(&response)
}

// CURL DETAIL POKEMON
func (h handler) GetListPokemons(c *fiber.Ctx) error {
	id := c.Params("id")
	client := &http.Client{}
	link := fmt.Sprintf("%s/%s", "https://pokeapi.co/api/v2/pokemon", id)
	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyText, err := ioutil.ReadAll(resp.Body)
	// fmt.Printf("%s\n", bodyText)
	if err != nil {
		log.Fatal(err)
	}

	var jsonData = []byte(bodyText)
	var detail response.DetailPokemonAPI

	err = json.Unmarshal(jsonData, &detail)
	if err != nil {
		fmt.Println(err.Error())
	}

	var res response.PokemonImages
	copier.Copy(&res, &detail.Sprites)

	data := response.DetailPokemonAPI{
		Name:    detail.Name,
		Sprites: detail.Sprites,
	}
	response := response.DefaultResponse{
		Code:    200,
		Message: "Successfully get the Pokemon List",
		Data:    data,
	}

	return c.Status(fiber.StatusOK).JSON(&response)
}

// func (h handler) CreateAPIPokemons(c *fiber.Ctx) error {

// 	var pokemons []models.Pokemon
// 	// var res = response.PokemonListAPI{}
// 	// var data = ""
// 	// var data models.Pokemon

// 	if isExists := h.DB.Find(&pokemons); isExists.RowsAffected < 1 {
// 		for i := 1; i <= 3; i++ {
// 			data := h.GetAPIListPokemon(i)
// 			// fmt.Println(data.Sprites.FrontDefault)
// 			h.DB.Create(&data)
// 		}
// 	}

// 	response := response.DefaultResponse{
// 		Code:    200,
// 		Message: "Successfully create Pokemon Data",
// 	}

// 	return c.Status(fiber.StatusOK).JSON(&response)
// }
