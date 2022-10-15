package handler

//
import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/Gentleelephant/vue-project-backend/common"

	"github.com/Gentleelephant/vue-project-backend/config"
	"github.com/Gentleelephant/vue-project-backend/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var DB *gorm.DB

func init() {
	config.Filepath = "../config/config.yaml"
	config.Initial()
	DB = config.DB
}

type Router struct {
	Path string `json:"path"`

	Name string `json:"name"`

	Component string `json:"component"`

	Children []Router `json:"children"`
}

//var answer []Router

func TestGetCompleteAccountSQL(t *testing.T) {

	var account model.Account
	DB.Model(&model.Account{}).Preload("AccountInfo").Preload("Roles").Where("aid = ?", 1).First(&account)
	fmt.Println(account)

	marshal, err := json.Marshal(account)
	if err != nil {
		return
	}
	fmt.Println(string(marshal))

}

//func recursive(id int, db *gorm.DB) {
//	var temp []model.Menu
//	db.Where("parent_id = ?", id).Find(&temp)
//}

func TestCreateAccountWithDefaultRoleSQL(t *testing.T) {
	account, err := AddAccount(DB, &model.Account{
		Userid:   "u-user1",
		Username: "test-user1",
		Password: "password",
		Email:    "user4@qq.com",
		AccountInfo: model.AccountInfo{
			Icon: "test-user1",
		},
	})
	if err != nil {
		return
	}
	fmt.Println(account)
}

func TestDeleteAccountSQL(t *testing.T) {
	var account model.Account
	err := DB.Model(model.Account{}).Where("aid = ?", 3).Find(&account).Updates(model.Account{
		Userid:   fmt.Sprintf("%s%s", common.DeletePrefix(), account.Userid),
		Username: fmt.Sprintf("%s%s", common.DeletePrefix(), account.Username),
		Status:   common.Delete,
	}).Error
	if err != nil {
		return
	}
}

func TestDeleteAccountFunc(t *testing.T) {
	account, err := DeleteAccount(DB, &common.QueryAccountParams{
		Aid: 4,
	})
	if err != nil {
		return
	}
	fmt.Println(account)
}
func TestSelect(t *testing.T) {

	var roles []model.Role

	DB.Debug().Model(&model.Role{}).Preload("Menus").Where("rid = ?", 3).First(&roles)

	bytes, err := json.Marshal(roles)

	if err != nil {
		return
	}

	fmt.Println(string(bytes))

}

func TestRecursive(t *testing.T) {

	//menus := recursive(0, 0, DB)/
	//
	//marshal, err := json.Marshal(menus)
	//if err != nil {
	//	return
	//}
	//fmt.Println(string(marshal))

}

func TestCreateAccountCompleteSQL(t *testing.T) {

	var account = model.Account{
		Userid:   "user-id2",
		Username: "user-name2",
		Password: "password",
		Gender:   "男",
		Email:    "xxx@qq.com",
		Status:   common.Enable,
		AccountInfo: model.AccountInfo{
			Icon: "user-icon2",
		},
		Roles: []*model.Role{
			{
				RoleName:   "test-role2",
				RoleDesc:   "test-role2",
				RoleStatus: common.Enable,
			},
		},
	}

	err := DB.Create(&account).Error
	if err != nil {
		t.Fatal(err)
	}

}

//

//func TestGetUserRole(t *testing.T) {
//
//
//	role, err := GetRoleByAccountId(DB, 1, common.Enable)
//	if err != nil {
//		return
//	}
//
//	marshal, err := json.Marshal(role)
//	if err != nil {
//		return
//	}
//	fmt.Println(string(marshal))
//
//}

func TestGetMenusByRole(t *testing.T) {

	var roles []model.Role
	err := DB.Debug().Preload("Menus").Find(&roles).Error
	if err != nil {
		return
	}
	marshal, err := json.Marshal(roles)
	if err != nil {
		return
	}
	fmt.Println(string(marshal))

}

func TestGetAccountWithRole(t *testing.T) {

	var account []model.Account
	err := DB.Debug().Where(&model.Account{Aid: 1}).Preload("Roles").Find(&account).Error
	if err != nil {
		return
	}
	marshal, err := json.Marshal(account)
	if err != nil {
		return
	}
	fmt.Println(string(marshal))

}

func TestGetRoleByAccountId(t *testing.T) {

	var role []*model.Role
	var account model.Account
	err := DB.Debug().Where(&model.Account{
		Aid: 1,
	}).Preload("Roles").Find(&account).Error
	if err != nil {
		return
	}
	role = account.Roles
	marshal, err := json.Marshal(role)
	if err != nil {
		return
	}
	fmt.Println(string(marshal))

}

func TestAssign(t *testing.T) {

	var account model.Account
	err := DB.Debug().Where(&model.Account{Aid: 1}).Preload("Roles.Menus").Preload(clause.Associations).Find(&account).Error
	if err != nil {
		return
	}
	marshal, err := json.Marshal(account)
	if err != nil {
		return
	}
	fmt.Println(string(marshal))

	var menu []*model.Menu
	for i := range account.Roles {
		menu = append(menu, account.Roles[i].Menus...)
	}

	bytes, err := json.Marshal(menu)
	if err != nil {
		return
	}
	fmt.Println(string(bytes))

}

//func TestGetMenusByAccountId(t *testing.T) {
//	menus, err := GetMenuByAccountId(DB, 1)
//	if err != nil {
//		return
//	}
//	marshal, err := json.Marshal(menus)
//	if err != nil {
//		return
//	}
//	fmt.Println(string(marshal))
//	tree := model.ArrayToTree(menus)
//	marshal, err = json.Marshal(tree)
//	if err != nil {
//		return
//	}
//	fmt.Println(string(marshal))
//}
