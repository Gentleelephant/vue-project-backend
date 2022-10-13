package service

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/Gentleelephant/vue-project-backend/handler"
	"gorm.io/gorm"
)

// CheckPassword checks if the password is correct.
func CheckPassword(db *gorm.DB, username string, password string) (bool, error) {
	account, err := handler.FindAccountByUsername(db, username)
	if err != nil {
		return false, err
	}
	if account.Password == password {
		return true, err
	}
	return false, err
}

func sha256encode(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	sum := hash.Sum(nil)
	pd := hex.EncodeToString(sum)
	return pd
}
