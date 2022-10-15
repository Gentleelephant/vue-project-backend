package handler

import (
	"encoding/json"
	"testing"

	"github.com/Gentleelephant/vue-project-backend/common"
	"github.com/Gentleelephant/vue-project-backend/model"
)

func TestGetMenuByAccountSQL(t *testing.T) {

	var params = &common.QueryAccountParams{
		Aid: 1,
	}
	var account model.Account
	err := DB.Where(&model.Account{
		Aid:      params.Aid,
		Userid:   params.Userid,
		Username: params.Username,
	}).Preload("Roles.Menus", "menu_status = ?", common.Enable).Preload("Roles", "role_status = ?", common.Enable).Find(&account).Error

	if err != nil {
		t.Error(err)
	}

	var menu []*model.Menu

	for i := range account.Roles {
		menu = append(menu, account.Roles[i].Menus...)
	}

	// 去重
	menu = removeDuplication(menu)

	marshal, err := json.Marshal(menu)
	if err != nil {
		return
	}
	t.Log(string(marshal))

}

func TestGetMenuByAccount(t *testing.T) {
	var params = &common.QueryAccountParams{
		Aid: 1,
	}
	menu, err := GetMenuByAccount(DB, params)
	if err != nil {
		return
	}
	removeDuplication(menu)
	router := model.ArrayToTree(menu)
	marshal, err := json.Marshal(router)
	if err != nil {
		return
	}
	t.Log(string(marshal))
}

func TestAddMenuWithRoleSQL(t *testing.T) {

	var menus = []*model.Menu{
		{
			MenuName:   "test-2",
			ParentId:   0,
			Path:       "login",
			Component:  "login.vue",
			MenuStatus: common.Enable,
			Roles:      []*model.Role{{Rid: 1}},
			Children:   nil,
		},
	}

	DB.Model(&model.Menu{}).Create(menus)

}
