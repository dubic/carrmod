package services

import (
	"carrmod/backend/domain/dto"
	"carrmod/backend/domain/models"
	"log"
)

type UserService struct {
	userRepo *models.UserRepo
}

func NewUserService(userRepo *models.UserRepo) *UserService {
	return &UserService{userRepo}
}

func (userService *UserService) CreateAccount(userCreationRequest dto.UserCreationRequest) error {
	hashedPassword := HashPassword(userCreationRequest.Password)

	user := models.NewUser(userCreationRequest, hashedPassword)
	err := userService.userRepo.SaveNewUser(user)
	log.Println("created and saved new user: ", user)
	log.Println("Send mail")
	return err
}
