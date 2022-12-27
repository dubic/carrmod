package api

import (
	"carrmod/backend/domain/dto"
	"carrmod/backend/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var userService *services.UserService

func UserRoutes(router *gin.Engine, us *services.UserService) {
	userService = us
	root := router.Group("/users")
	{
		root.POST("/createAccount", createAccount)
	}
}

// create a new profile using email as username and password
func createAccount(c *gin.Context) {
	var userCreationRequest dto.UserCreationRequest
	c.ShouldBindJSON(&userCreationRequest)
	log.Println("Received account creation request : ", userCreationRequest)

	userCreationRequest.Create()
	userService.CreateAccount(userCreationRequest)
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"msg":    "successfully created account",
	})
}
