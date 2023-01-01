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
	key := os.Getenv("JWT_KEY")
	// log.Println(key)
	// log.Println(key.([]byte))
	tokenString, err := token.SignedString([]byte(key))
	return tokenString, err
}

// parse the jwt token and returns the username if valid
func ParseJwt(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_KEY")), nil
	})

	if token == nil {
		return "", fmt.Errorf("invalid token: %s", tokenString)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["principal"].(string), nil
	}
	return "", err
}
