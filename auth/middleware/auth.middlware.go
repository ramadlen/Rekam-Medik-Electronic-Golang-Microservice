package middleware

import (
	"biodata/utils"

	"github.com/gofiber/fiber/v2"
)
func AuthMiddleware(ctx *fiber.Ctx) error {
token := ctx.Get("x-token")
if token == "" {
	return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"messsage": "unathenthicated",
	})
}
_, err  := utils.VerifyToken(token)
if err != nil{
	return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"message": "unauthenticated",
		})
	}
	return ctx.Next()
}


func PermissionCreate(ctx *fiber.Ctx) error {
	return ctx.Next()
}