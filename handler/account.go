package handler

import (
	"fmt"

	"github.com/Gentleelephant/vue-project-backend/common"
	"github.com/Gentleelephant/vue-project-backend/model"
	"gorm.io/gorm"
)

func GetAccount(db *gorm.DB, params *common.QueryAccountParams) (*model.Account, error) {
	var account model.Account
	err := db.Where(model.Account{
		Aid:      params.Aid,
		Userid:   params.Userid,
		Username: params.Username,
	}).First(&account).Error
	if err != nil {
		return nil, err
	}
	return &account, err
}

func AddAccount(db *gorm.DB, account *model.Account) (*model.Account, error) {
	var role = model.DefaultRole
	if account.Roles == nil {
		account.Roles = append(account.Roles, &role)
	}
	account.Status = common.Enable
	err := db.Create(account).Error
	if err != nil {
		return nil, err
	}
	return account, nil
}

func DeleteAccount(db *gorm.DB, params *common.QueryAccountParams) (*model.Account, error) {
	var account model.Account
	err := db.Model(model.Account{}).Where(model.Account{
		Aid:      params.Aid,
		Username: params.Username,
		Userid:   params.Userid,
	}).Find(&account).Updates(model.Account{
		Userid:   fmt.Sprintf("%s%s", common.DeletePrefix(), account.Userid),
		Username: fmt.Sprintf("%s%s", common.DeletePrefix(), account.Username),
		Status:   common.Delete,
	}).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func UpdateAccount(db *gorm.DB, account *model.Account) (*model.Account, error) {
	err := db.Model(account).Updates(account).Error
	if err != nil {
		return nil, err
	}
	return account, nil
}

func IsAccountExist(db *gorm.DB, params *common.QueryAccountParams) (bool, error) {
	tx := db.Where(model.Account{
		Aid:      params.Aid,
		Username: params.Username,
		Userid:   params.Userid,
		Status:   common.Enable,
	}).First(model.Account{})
	err := tx.Error
	if err != nil {
		return false, err
	}
	if tx.RowsAffected > 0 {
		return true, nil
	}
	return false, nil
}
