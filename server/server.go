package server

import (
	"konecta/config"
	"konecta/feature"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Server struct {
	cfg     config.Config
	feature feature.Feature
}

func NewServer(cfg config.Config, feature feature.Feature) Server {
	return Server{
		cfg:     cfg,
		feature: feature,
	}
}

func (s *Server) Run() error {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "HEAD,GET,POST,PUT,DELETE,PATCH",
		AllowCredentials: false,
	}))
	s.RegisterRoute(app)
	return app.Listen(s.cfg.Host)

}
