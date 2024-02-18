package routes

import (
	"goapiauth/models"
	"goapiauth/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch events.",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Events",
		"events":  events,
	})
}

func getEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event id.",
		})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not find the event.",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"event": event,
	})
}

func createEvents(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorize access",
		})
		return
	}

	err := utils.VerifyToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorize access. Invalid token",
		})
		return
	}

	var event models.Event
	err = c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request",
		})
		return
	}
	// event.ID = 1
	event.UserID = 1234
	err = event.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create event.",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Event created",
		"event":   event,
	})
}

func updateEvents(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event id.",
		})
		return
	}
	_, err = models.GetEventById(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch the event.",
		})
		return
	}
	var updatedEvent models.Event
	err = c.ShouldBindJSON(&updatedEvent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request",
		})
		return
	}
	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not update event",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Event updated successfully.",
	})
}
