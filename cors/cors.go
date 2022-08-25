package cors

import (
	"github.com/fumeapp/tonic/setting"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func New() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{setting.Core.WebURL},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           0,
	})
}
