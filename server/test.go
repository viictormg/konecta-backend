package server

import "github.com/gofiber/fiber/v2"

func (s *Server) TestHadler(ctx *fiber.Ctx) error {

	// statusCode := fiber.StatusOK
	// var account models.SingInRequest
	// err := ctx.BodyParser(&account)

	// if err != nil {
	// 	return err
	// }
	// response := models.ResponseStandar{
	// 	Success: true,
	// 	Message: "cuenta creada con exito",
	// }

	// err = s.feature.SignIn(account)

	// if err != nil {
	// 	response.Success = false
	// 	response.Message = "error creando cuenta"
	// 	response.Error = err

	// 	statusCode = fiber.StatusInternalServerError
	// }

	// return ctx.Status(statusCode).JSON(response)
	return nil
}
