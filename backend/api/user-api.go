package api

import (
	"carrmod/backend/domain/dto"
	"carrmod/backend/services"
	"log"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	svc *services.UserService
}

func RegisterUserController(router *gin.Engine, us *services.UserService) {
	uc := UserController{us}
	root := router.Group("/users")
	{
		root.POST("/v1/createAccount", uc.createAccount)
	}
}

// create a new profile using email as username and password
func (uc UserController) createAccount(c *gin.Context) {
	var userCreationRequest dto.UserCreationRequest
	c.ShouldBindJSON(&userCreationRequest)
	log.Println("Received account creation request : ", userCreationRequest.Print())

	userCreationRequest.Create()
	userCreationResponse, err := uc.svc.CreateAccount(userCreationRequest)
	Respond(c, userCreationResponse, err)
}
