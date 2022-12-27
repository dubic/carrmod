package models

import (
	"carrmod/backend/domain/dto"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Name     string `bson:"name"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
}

type UserRepo struct {
	users *mongo.Collection
}

func NewUser(userDto dto.UserCreationRequest, hashedPassword string) User {
	return User{Name: userDto.Name, Email: userDto.Email, Password: hashedPassword}
}

func NewUserRepo(col *mongo.Collection) *UserRepo {
	return &UserRepo{col}
}

func (repo UserRepo) SaveNewUser(user User) error {
	_, err := repo.users.InsertOne(context.TODO(), user)
	return err
}
