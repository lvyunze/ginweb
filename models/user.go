package models

import (
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
