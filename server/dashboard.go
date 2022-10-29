package server

import (
	"konecta/model"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (s Server) GetDashboard(ctx *fiber.Ctx) error {
	reponse, err := s.feature.GetDashboard()

	if err != nil {
		return err
	}

	response := model.ReponseStandar{
		Status:  true,
		Message: "Dashboard",
		Data:    reponse,
	}

	return ctx.Status(http.StatusOK).JSON(response)
}
