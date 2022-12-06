package services

import (
	"carrmod/backend/domain/dto"
	"log"
)

type UserService struct {
}

func CreateAccount(userCreationRequest dto.UserCreationRequest) error {
	log.Println("Hash password")
	log.Println("create and save user")
	log.Println("Send mail")
	return nil
}
