package routes

import (
	"goapiauth/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/event/:id", getEvent)
	server.POST("/signup", signup)
	server.POST("/login", login)

	// PROTECTED (AUTH) ROUTES
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvents)
	authenticated.PUT("/event/:id", updateEvents)
}
