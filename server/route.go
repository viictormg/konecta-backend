package server

import "github.com/gofiber/fiber/v2"

func (s *Server) RegisterRoute(app *fiber.App) {
	routes := app.Group("/api")

	routes.Post("/test", s.TestHadler)
}
