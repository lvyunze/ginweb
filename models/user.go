package models

import (
	"ginweb/ext"
	"ginweb/utils/common"
	"gorm.io/gorm"
)

// User data model
type User struct {
	gorm.Model   `json:"gorm_._model,omitempty"`
	Username     string `json:"username,omitempty"`
	DisplayName  string `json:"display_name,omitempty"`
	PasswordHash string `json:"password_hash,omitempty"`
}

// Serialize serializes user data
func (u *User) Serialize() common.JSON {
	return common.JSON{
		"id":           u.ID,
		"username":     u.Username,
		"display_name": u.DisplayName,
	}
}

func GetUser() (userList []*User, err error) {
	if err = ext.DB.Find(&userList).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateUser() (userList []*User, err error) {
	if err = ext.DB.Find(&userList).Error; err != nil {
		return nil, err
	}
	return
}

func GetAUser(id string) (user *User, err error) {
	user = new(User)
	if err = ext.DB.Where("id=?", id).First(user).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateAUser(user *User) (err error) {
	err = ext.DB.Save(user).Error
	return
}

func DeleteAUser(id string) (err error) {
	err = ext.DB.Where("id=?", id).Delete(&User{}).Error
	return
}
