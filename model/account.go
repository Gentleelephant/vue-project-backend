package model

// Account is a struct that represents a user account.
type Account struct {
	Aid         int         `json:"aid" gorm:"primary_key;autoIncrement"`
	Userid      string      `json:"userid" gorm:"type:varchar(36);unique;not null;uniqueIndex"`
	Username    string      `json:"username" gorm:"type:varchar(36);not null;unique;uniqueIndex"`
	Password    string      `json:"-" gorm:"type:varchar(256);not null"`
	Gender      string      `json:"gender"   gorm:"type:varchar(10);not null;default:'男'"`
	Email       string      `json:"email" gorm:"type:varchar(36)"`
	Status      byte        `json:"status" gorm:"type:tinyint(1);not null;default:0"`
	CreateAt    int         `json:"create_at" gorm:"autoCreateTime"`
	UpdateAt    int         `json:"update_at" gorm:"autoUpdateTime"`
	AccountInfo AccountInfo `json:"accountInfo" gorm:"foreignKey:Aid"`    // 以另一个关系的外键作主关键字的表被称为主表,所以account是主表，account_info是从表
	Roles       []*Role     `json:"roles" gorm:"many2many:account_role;"` // many2many:account_role; 为中间表名
}

func (a *Account) IgnorePassword() *Account {
	a.Password = ""
	return a
}

// RegisterAccount RegisterAccount.
type RegisterAccount struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
}

type UserData struct {
	SessionId string  `json:"sessionId"`
	Account   Account `json:"account"`
}
