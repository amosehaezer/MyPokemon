package main

import (
	"log"
	"phincon/pkg/common/config"
	"phincon/pkg/common/db"
	"phincon/pkg/pokemons"
	"phincon/pkg/users"

	"github.com/gofiber/fiber/v2"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(&c)
	app := fiber.New()

	users.RegisterRoutes(app, h)
	pokemons.RegisterRoutes(app, h)

	app.Listen(c.Port)
}
