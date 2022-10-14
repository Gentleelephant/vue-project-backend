package config

import (
	"encoding/json"
	"fmt"
	"testing"

	"gorm.io/gorm/clause"

	"github.com/Gentleelephant/vue-project-backend/handler"

	"gorm.io/gorm"

	"github.com/Gentleelephant/vue-project-backend/model"
)

type Router struct {
	Path string `json:"path"`

	Name string `json:"name"`

	Component string `json:"component"`

	Children []Router `json:"children"`
}

//var answer []Router

func TestConfig(t *testing.T) {
	ProjectConfig = loadConfig("config.yaml")
	DB = initDB()
	createTable(DB)
	//DB.Create(&model.Account{
	//	Userid:   "u-12345",
	//	Username: "account",
	//	Password: "password",
	//	Email:    "1132960613@qq.com",
	//	Status:   0,
	//	AccountInfo: model.AccountInfo{
	//		Icon: "test-icon",
	//	},
	//	Roles: []model.Role{
	//		{
	//			RoleName:   "role1",
	//			RoleDesc:   "role test1",
	//			RoleStatus: 0,
	//		},
	//		{
	//			RoleName:   "role2",
	//			RoleDesc:   "role test2",
	//			RoleStatus: 0,
	//		},
	//	},
	//})
	//DB.Create(&model.Menu{
	//	MenuName:  "menu test",
	//	Level:     0,
	//	ParentId:  0,
	//	Path:      "/home",
	//	Component: "xxx.vue",
	//	Roles: []model.Role{
	//		{
	//			Rid: 0,
	//		},
	//	},
	//})

	var account model.Account
	DB.Model(&model.Account{}).Preload("AccountInfo").Preload("Roles").Where("aid = ?", 1).First(&account)
	fmt.Println(account)

	marshal, err := json.Marshal(account)
	if err != nil {
		return
	}
	fmt.Println(string(marshal))

	var menu []model.Menu

	DB.Model(&model.Menu{}).Preload("Roles").Find(&menu)

	bytes, err := json.Marshal(menu)
	if err != nil {
		return
	}
	fmt.Println(string(bytes))

}

//func recursive(id int, db *gorm.DB) {
//	var temp []model.Menu
//	db.Where("parent_id = ?", id).Find(&temp)
//}

func TestCreate(t *testing.T) {
	ProjectConfig = loadConfig("config.yaml")
	DB = initDB()
	createTable(DB)
	//DB.Create(&model.Account{
	//	Userid:   "u-12345",
	//	Username: "account",
	//	Password: "password",
	//	Email:    "1132960613@qq.com",
	//	Status:   0,
	//	AccountInfo: model.AccountInfo{
	//		Icon: "test-icon",
	//	},
	//	Roles: []model.Role{
	//		{
	//			RoleName:   "role1",
	//			RoleDesc:   "role test1",
	//			RoleStatus: 0,
	//		},
	//		{
	//			RoleName:   "role2",
	//			RoleDesc:   "role test2",
	//			RoleStatus: 0,
	//		},
	//	},
	//})
}

func TestSelect(t *testing.T) {

	ProjectConfig = loadConfig("config.yaml")
	DB = initDB()
	createTable(DB)

	var roles []model.Role

	DB.Debug().Model(&model.Role{}).Preload("Menus").Where("rid = ?", 2).First(&roles)

	bytes, err := json.Marshal(roles)

	if err != nil {
		return
	}

	fmt.Println(string(bytes))

}

func TestRecursive(t *testing.T) {

	ProjectConfig = loadConfig("config.yaml")
	DB = initDB()
	createTable(DB)

	menus := recursive(0, 0, DB)

	marshal, err := json.Marshal(menus)
	if err != nil {
		return
	}
	fmt.Println(string(marshal))

}

func recursive(level, parent int, db *gorm.DB) []model.Menu {
	var temp []model.Menu
	db.Where("level = ? and parent_id = ?", level, parent).Find(&temp)
	for i := range temp {
		temp[i].Children = recursive(level+1, temp[i].Mid, db)
	}
	return temp
}

func TestGetUserRole(t *testing.T) {

	ProjectConfig = loadConfig("config.yaml")
	DB = initDB()
	createTable(DB)

	role, err := handler.GetRoleByAccountId(DB, 1)
	if err != nil {
		return
	}

	marshal, err := json.Marshal(role)
	if err != nil {
		return
	}
	fmt.Println(string(marshal))

}

func TestGetMenusByRole(t *testing.T) {

	ProjectConfig = loadConfig("config.yaml")
	DB = initDB()
	createTable(DB)

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

	ProjectConfig = loadConfig("config.yaml")
	DB = initDB()
	createTable(DB)

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

	ProjectConfig = loadConfig("config.yaml")
	DB = initDB()
	createTable(DB)

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

	ProjectConfig = loadConfig("config.yaml")
	DB = initDB()
	createTable(DB)

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

func TestGetMenusByaccountId(t *testing.T) {
	ProjectConfig = loadConfig("config.yaml")
	DB = initDB()
	createTable(DB)
	menus, err := handler.GetMenuByAccountId(DB, 1)
	if err != nil {
		return
	}
	marshal, err := json.Marshal(menus)
	if err != nil {
		return
	}
	fmt.Println(string(marshal))
	tree := model.ArrayToTree(menus)
	marshal, err = json.Marshal(tree)
	if err != nil {
		return
	}
	fmt.Println(string(marshal))
}
