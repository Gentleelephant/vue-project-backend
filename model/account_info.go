package model

type AccountInfo struct {
	InfoId int `json:"infoId" gorm:"primary_key;autoIncrement"`

	Aid int `json:"-" gorm:"type:int(11);not null"`

	Icon string `json:"icon" gorm:"type:varchar(255);not null"`
}
