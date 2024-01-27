package main

import (
	"goapiauth/db"
	"goapiauth/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	r := gin.Default()
	r.GET("/events", getEvents)
	r.POST("/events", createEvents)
	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func getEvents(c *gin.Context) {
	events := models.GetAllEvents()
	c.JSON(http.StatusOK, gin.H{
		"message": "Events",
		"events":  events,
	})
}

func createEvents(c *gin.Context) {
	var event models.Event
	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request",
		})
		return
	}
	event.ID = 1
	event.UserID = 1234
	event.Save()
	c.JSON(http.StatusCreated, gin.H{
		"message": "Event created",
		"event":   event,
	})
}
