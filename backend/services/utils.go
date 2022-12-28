package services

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func HashPassword(plain string, salt string) string {
	hash := sha256.Sum256([]byte(fmt.Sprintf("%scarrmod%s", plain, salt)))
	return hex.EncodeToString(hash[:])
}
