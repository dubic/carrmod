package models

import (
	"carrmod/backend/domain/dto"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Base     BaseModel `bson:"inline"`
	Name     string    `bson:"name"`
	Email    string    `bson:"email"`
	Password string    `bson:"password"`
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

func (repo UserRepo) FindUserByEmail(email string) (User, error) {
	var user User
	err := repo.users.FindOne(context.TODO(), bson.D{{Key: "email", Value: email}}).Decode(&user)
	return user, err
}

func (repo UserRepo) UserExists(email string) (bool, error) {
	count, err := repo.users.CountDocuments(context.TODO(), bson.D{{Key: "email", Value: email}})
	if count > 0 {
		return true, err
	}
	return false, err
}
