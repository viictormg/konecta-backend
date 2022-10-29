package server

import "github.com/gofiber/fiber/v2"

func (s *Server) RegisterRoute(app *fiber.App) {
	routes := app.Group("/api")

	routes.Post("/test", s.TestHadler)
	routes.Get("/auth", s.AuthKonecta)
	routes.Get("/dashboard", s.GetDashboard)
	routes.Get("/dashboard-local/:id", s.GetInfoDashboardDBByID)
	routes.Get("/agents", s.GetAgents)

}
