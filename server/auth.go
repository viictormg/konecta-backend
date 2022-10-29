package server

import (
	"konecta/model"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (s Server) AuthKonecta(ctx *fiber.Ctx) error {
	token, err := s.feature.Auth()

	if err != nil {
		return err
	}

	response := model.ReponseStandar{
		Status:  true,
		Message: "hola",
		Data:    token,
	}

	return ctx.Status(http.StatusOK).JSON(response)
}
