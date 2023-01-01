package models

import (
	"carrmod/backend/domain/dto"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Base          TimeModel `bson:"inline"`
	Name          string    `bson:"name"`
	Email         string    `bson:"email"`
	Password      string    `bson:"password"`
	EmailVerified bool      `bson:"emailVerified"`
}

type UserRepo struct {
	users *mongo.Collection
}

type Session struct {
	User       string
	ExpiresAt  time.Time `bson:"expires_at"`
	AccessedAt time.Time `bson:"accessed_at"`
}

type SessionRepo struct {
	sessions *mongo.Collection
}

func NewUser(userDto dto.UserCreationRequest, hashedPassword string) User {
	return User{Name: userDto.Name, Email: userDto.Email, Password: hashedPassword, EmailVerified: false,
		Base: TimeModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}}
}

func NewUserRepo(col *mongo.Collection) *UserRepo {
	return &UserRepo{col}
}

func NewSessionRepo(col *mongo.Collection) *SessionRepo {
	return &SessionRepo{col}
}

func (u User) Print() string {
	return fmt.Sprintf("Name: %s, Email: %s", u.Name, u.Email)
}

// insert user in db
func (repo UserRepo) SaveNewUser(user User) error {
	_, err := repo.users.InsertOne(context.TODO(), user)
	return err
}

// find a single user by email as filter
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

// find a single user by email and as filter
func (repo UserRepo) FindUserByEmailAndPassword(email string, password string) (User, error) {
	var user User
	err := repo.users.FindOne(context.TODO(), bson.D{
		{Key: "email", Value: email},
		{Key: "password", Value: password}}).Decode(&user)
	return user, err
}

// -----------------------SESSION REPO---------------------

// insert session in db
func (repo SessionRepo) SaveNewSession(session Session) error {
	_, err := repo.sessions.InsertOne(context.TODO(), session)
	return err
}

// delete session in db
func (repo SessionRepo) DeleteSession(session Session) error {
	_, err := repo.sessions.DeleteOne(context.TODO(), bson.D{
		{Key: "user", Value: session.User},
	})
	return err
}

// delete all sessions of a user with email. user can have multiple sessions
func (repo SessionRepo) DeleteSessions(email string) int64 {
	result, err := repo.sessions.DeleteMany(context.TODO(), bson.D{
		{Key: "user", Value: email},
	})
	if err != nil {
		log.Printf("error deleting [%s] sessions: %s", email, err)
		return 0
	}
	return result.DeletedCount
}
