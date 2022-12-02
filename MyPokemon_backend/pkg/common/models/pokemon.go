package models

import "gorm.io/gorm"

type PokemonType struct {
	gorm.Model
	Name         string `json:"name"`
	Url          string `json:"url"`
	PokemonRefer uint
}
type PokemonMoves struct {
	gorm.Model
	Name         string `json:"name"`
	Url          string `json:"url"`
	PokemonRefer uint
}
type Sprites struct {
	gorm.Model
	FrontDefault string `json:"front_default"`
	BackDefault  string `json:"back_default"`
	FrontShiny   string `json:"front_shiny"`
	BackShiny    string `json:"back_shiny"`
	SpriteRefer  uint
}

// POKEMON Home + Detail
type Pokemon struct {
	gorm.Model
	Name         string         `json:"name"`
	Sprites      []Sprites      `json:"sprites" gorm:"foreignKey:SpriteRefer"`
	PokemonType  []PokemonType  `gorm:"foreignKey:PokemonRefer"`
	PokemonMoves []PokemonMoves `gorm:"foreignKey:PokemonRefer"`
}

// API LIST POKEMON
type ListPokemon struct {
	Count    int             `json:"count"`
	Next     string          `json:"next"`
	Previous string          `json:"previous"`
	Result   []ResultPokemon `json:"results"`
}
type ResultPokemon struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

// MY POKEMON
type MyPokemon struct {
	gorm.Model
	PokemonID   uint   `json:"pokemon_id"`
	PokemonName string `json:"pokemon_name"`
}
