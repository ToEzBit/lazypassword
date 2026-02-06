package crypto

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"

	"golang.org/x/crypto/pbkdf2"
)

const SaltSize = 32
const KeySize = 32
const Iterations = 100000

func GenerateSalt() (string, error) {
	salt := make([]byte, SaltSize)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(salt), nil
}

// HashPassword creates a PBKDF2 hash of the password using the provided salt
func HashPassword(password, salt string) string {
	saltBytes, _ := base64.StdEncoding.DecodeString(salt)
	hash := pbkdf2.Key([]byte(password), saltBytes, Iterations, KeySize, sha256.New)
	return base64.StdEncoding.EncodeToString(hash)
}

// VerifyPassword checks if the password matches the stored hash
func VerifyPassword(password, salt, storedHash string) bool {
	computedHash := HashPassword(password, salt)
	return computedHash == storedHash
}
