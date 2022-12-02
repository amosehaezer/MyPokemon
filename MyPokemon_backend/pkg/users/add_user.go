package users

import (
	"phincon/pkg/common/models"

	"github.com/gofiber/fiber/v2"
)

type AddUserRequestBody struct {
	Name string `json:"name"`
}

func (h handler) AddUser(c *fiber.Ctx) error {
	body := AddUserRequestBody{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var user models.User

	user.Name = body.Name

	if result := h.DB.Create(&user); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&user)
}
