package server

import (
	"konecta/model"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (s Server) GetDashboard(ctx *fiber.Ctx) error {
	reponse, err := s.feature.GetDashboardKonecta()

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

func (s Server) GetInfoDashboardDBByID(ctx *fiber.Ctx) error {

	id := ctx.Params("id", "")

	reponse, err := s.feature.GetInfoDashboardDBByID(id)

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

func (s Server) GetAgents(ctx *fiber.Ctx) error {
	response, err := s.feature.GetAgents()

	if err != nil {
		return err
	}

	responseHandler := model.ReponseStandar{
		Status:  true,
		Message: "agentes encontrados",
		Data:    response,
	}

	return ctx.Status(http.StatusOK).JSON(responseHandler)
}
