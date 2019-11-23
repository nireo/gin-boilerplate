package models

import (
	"github.com/jinzhu/gorm"
	"github.com/nireo/gin-boilerplate/lib/common"
)

// User data alias
type User struct {
	gorm.Model
	Username string
	Password string
}

// Serialize user data
func (u *User) Serialize() common.JSON {
	return common.JSON{
		"id":       u.ID,
		"username": u.Username,
	}
}

// Read User Data
func (u *User) Read(m common.JSON) {
	u.ID = uint(m["id"].(float64))
	u.Username = m["username"].(string)
}
