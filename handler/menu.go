package handler

import (
	"github.com/Gentleelephant/vue-project-backend/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetMenuByAccountId(db *gorm.DB, aid int) ([]*model.Menu, error) {
	var account model.Account
	err := db.Debug().Where(&model.Account{Aid: aid}).Preload("Roles.Menus").Preload(clause.Associations).Find(&account).Error
	if err != nil {
		return nil, err
	}
	var menu []*model.Menu
	for i := range account.Roles {
		menu = append(menu, account.Roles[i].Menus...)
	}
	// 去重
	menu = removeDuplication(menu)
	return menu, err
}

func removeDuplication(arrays []*model.Menu) []*model.Menu {
	set := make(map[*model.Menu]struct{}, len(arrays))
	j := 0
	for _, v := range arrays {
		_, ok := set[v]
		if ok {
			continue
		}
		set[v] = struct{}{}
		arrays[j] = v
		j++
	}
	return arrays[:j]
}
