package models

import (
	"github.com/jinzhu/gorm"
)

// User model from database
type User struct {
	gorm.Model
	Email    string `gorm:"size:50;not null;unique"`
	Login    string `gorm:"size:50;not null;unique" valid:"required"`
	Password string `gorm:"size:255;not null" valid:"required"`
}
