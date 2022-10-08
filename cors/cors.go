package cors

import (
	"github.com/fumeapp/tonic/setting"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func New() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins:     setting.Core.WebURL,
		AllowCredentials: true,
		AllowHeaders:     "Origin, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, User-Agent",
		ExposeHeaders:    "Origin, User-Agent",
	})
}
