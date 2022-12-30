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
		root.POST("/v1/login", uc.login)
	}
}

// create a new profile using email as username and password
func (uc UserController) createAccount(c *gin.Context) {
	var userCreationRequest dto.UserCreationRequest
	c.ShouldBindJSON(&userCreationRequest)
	log.Println("Received account creation request : ", userCreationRequest.Print())

	ok := Handle(c, userCreationRequest.Create())
	if !ok {
		return
	}
	userCreationResponse, err := uc.svc.CreateAccount(userCreationRequest)
	Respond(c, userCreationResponse, err)
}

// Login with email
func (ctrl UserController) login(c *gin.Context) {
	var loginRequest dto.LoginRequest
	c.ShouldBindJSON(&loginRequest)
	log.Println("Received login request : ", loginRequest.Email)

	ok := Handle(c, loginRequest.Validate())
	if !ok {
		return
	}
	loginResponse := ctrl.svc.Login(loginRequest)
	Respond(c, loginResponse, nil)
}
