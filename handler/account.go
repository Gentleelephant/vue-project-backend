package handler

import (
	"github.com/Gentleelephant/vue-project-backend/model"
	"gorm.io/gorm"
)

func FindAccountByUsername(db *gorm.DB, username string) (*model.Account, error) {
	var account model.Account
	err := db.Where("username = ?", username).First(&account).Error
	if err != nil {
		return nil, err
	}
	return &account, err
}

func AddAccount(db *gorm.DB, account *model.Account) error {
	err := db.Create(account).Error
	if err != nil {
		return err
	}
	return nil
}

func AccountExistByUsername(db *gorm.DB, username string) (bool, error) {
	var account model.Account
	err := db.Where("username = ?", username).First(&account).Error
	if err != nil {
		return false, err
	}
	return true, err
}
