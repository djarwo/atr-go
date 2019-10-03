package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	ID                uint   `gorm:"primary_key:true;index;" json:"ID"`
	Code              string `json:"Code"`
	Username          string `gorm:"type:varchar(100);unique_index" form:"Username" json:"Username"`
	Password          string `json:"Password"`
	Name              string `json:"Name"`
	Email             string `gorm:"type:varchar(100);unique_index" form:"Email" json:"Email"`
	Phone             string `gorm:"type:varchar(100);unique_index" form:"Phone" json:"Phone"`
	Pin               string `json:"Pin"`
	LoginType         string `json:"LoginType"`
	UID               string `json:"UID"`
	Type              string `json:"Type"`
	Location          string `json:"Location"`
	Longitude         string `json:"Longitude"`
	Latitude          string `json:"Latitude"`
	SessionToken      string `json:"SessionToken"`
	DeviceToken       string `json:"DeviceToken"`
	NotificationToken string `json:"NotificationToken"`
	RefreshToken      string `json:"RefreshToken"`
	Description       string `json:"Description"`
}

func GetCodeUser(db *gorm.DB, prefix string) string {
	var user User
	var code string

	errCode := db.Last(&user).Error
	if errCode != nil {
		code = "0000000001"
	} else {
		code = user.Code
	}
	return code
}
