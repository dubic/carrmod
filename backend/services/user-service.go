package services

import (
	"carrmod/backend/domain/dto"
	"carrmod/backend/domain/models"
	"fmt"
	"log"
)

type UserService struct {
	userRepo *models.UserRepo
}

func NewUserService(userRepo *models.UserRepo) *UserService {
	return &UserService{userRepo}
}

// Creates an account. emails are unique.
// If an account with the same email exists, a soft error is returned.
func (svc *UserService) CreateAccount(userCreationRequest dto.UserCreationRequest) (dto.UserCreationResponse, error) {
	hashedPassword := HashPassword(userCreationRequest.Password)
	//check user exists
	exists, err := svc.userRepo.UserExists(userCreationRequest.Email)
	if err != nil {
		return dto.UserCreationResponse{}, err
	}
	if exists {
		return dto.UserCreationResponse{Created: false, Msg: fmt.Sprintf("User with email exists - %s", userCreationRequest.Email)}, err
	}
	//create user
	user := models.NewUser(userCreationRequest, hashedPassword)
	err = svc.userRepo.SaveNewUser(user)
	log.Println("created and saved new user: ", user)
	log.Println("Send mail")
	return dto.UserCreationResponse{VerificationMailSent: false, Created: true, Email: user.Email}, err
}
