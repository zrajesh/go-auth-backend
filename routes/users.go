package routes

import (
	"goapiauth/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request",
		})
		return
	}
	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not save user",
		})
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"message": "User signup successfully",
	})
}
