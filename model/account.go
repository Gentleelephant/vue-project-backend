package model

// Account is a struct that represents a user account.
type Account struct {
	Id       int    `json:"id" gorm:"primary_key;autoIncrement"`
	Username string `json:"username" gorm:"type:varchar(36);not null;unique"`
	Password string `json:"password" gorm:"type:varchar(256);not null"`
	Gender   string `json:"gender"   gorm:"type:varchar(10);not null;default:'男'"`
	Role     byte   `json:"role" gorm:"type:tinyint(1);default:0"`
	Email    string `json:"email" gorm:"type:varchar(36)"`
	Status   byte   `json:"status" gorm:"type:tinyint(1);not null;default:0"`
	CreateAt int    `json:"create_at" gorm:"autoCreateTime"`
	UpdateAt int    `json:"update_at" gorm:"autoUpdateTime"`
}
