package router

import (
	"net/http"
	"realtime-score/internal/middleware"

	"github.com/gin-gonic/gin"
)

func TestingRouter(app *gin.Engine) {
	authGroup := app.Group("/")
	authGroup.GET("/authenticate", middleware.AuthMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "You have access!",
			"user_id": c.GetString("user_id"),
			"role":    c.GetString("role"),
		})
	})
	authGroup.GET("/onlyadmin", middleware.AuthMiddleware(), middleware.OnlyAdmin(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "You have access!",
			"user_id": c.GetString("user_id"),
			"role":    c.GetString("role"),
		})
	})
}
