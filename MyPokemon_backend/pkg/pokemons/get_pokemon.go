package pokemons

// func (h handler) GetPokemon(c *fiber.Ctx) error {
// 	id := c.Params("id")

// 	var pokemon models.Pokemon

// 	if result := h.DB.First(&pokemon, id); result.Error != nil {
// 		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
// 	}

// 	return c.Status(fiber.StatusOK).JSON(&pokemon)
// }
