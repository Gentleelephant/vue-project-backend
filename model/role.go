package model

type Role struct {
	Rid int `json:"rid" gorm:"primary_key;autoIncrement"`

	RoleName string `json:"roleName" gorm:"type:varchar(50);not null"`

	RoleDesc string `json:"roleDesc" gorm:"type:varchar(50);not null"`

	RoleStatus byte `json:"roleStatus" gorm:"type:tinyint(1);not null;default:0"`

	RoleCreateAt int `json:"roleCreateAt" gorm:"autoCreateTime"`

	RoleUpdateAt int `json:"roleUpdateAt" gorm:"autoUpdateTime"`

	Menus []*Menu `json:"menuList" gorm:"many2many:menu_role;"`
}
