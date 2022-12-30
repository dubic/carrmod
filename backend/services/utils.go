package services

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

func HashPassword(plain string, salt string) string {
	hash := sha256.Sum256([]byte(fmt.Sprintf("%scarrmod%s", plain, salt)))
	return hex.EncodeToString(hash[:])
}

func GenerateJwt(principal string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"principal": principal,
	})

	// var key interface{}
	key := (os.Getenv("JWT_KEY"))
	// log.Println(key)
	// log.Println(key.([]byte))
	tokenString, err := token.SignedString([]byte(key))
	return tokenString, err
}
