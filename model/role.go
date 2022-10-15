package model

import "github.com/Gentleelephant/vue-project-backend/common"

type Role struct {
	Rid int `json:"rid" gorm:"primary_key;autoIncrement"`

	RoleName string `json:"roleName" gorm:"type:varchar(50);not nul;uniqueIndex"`

	RoleDesc string `json:"roleDesc" gorm:"type:varchar(50);not null"`

	RoleStatus byte `json:"roleStatus" gorm:"type:tinyint(1);not null;default:0"`

	RoleCreateAt int `json:"roleCreateAt" gorm:"autoCreateTime"`

	RoleUpdateAt int `json:"roleUpdateAt" gorm:"autoUpdateTime"`

	Accounts []*Account `json:"accountList" gorm:"many2many:account_role;"`

	Menus []*Menu `json:"menuList" gorm:"many2many:menu_role;"`
}

var DefaultRole = Role{
	Rid:        1,
	RoleName:   "default",
	RoleDesc:   "default role",
	RoleStatus: common.Enable,
}
