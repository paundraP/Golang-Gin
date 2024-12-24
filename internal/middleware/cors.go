package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupCors(app *gin.Engine) {
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // replace with your original fe url
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	app.OPTIONS("/*any", func(c *gin.Context) {
		c.AbortWithStatus(204)
	})
}
