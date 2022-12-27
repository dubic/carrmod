package services

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashPassword(plain string) string {
	hash := sha256.Sum256([]byte(plain))
	return hex.EncodeToString(hash[:])
}
