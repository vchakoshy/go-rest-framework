package models

import (
	"github.com/jinzhu/gorm"
)

// User model
type User struct {
	gorm.Model
	Username string `gorm:"index:'username'" json:"username"`
}
