package models

import "gorm.io/gorm"

type JWTUser struct {
	gorm.Model
	Name     string `json:"name" binding:"required"`
	Password string `gorm:"not null;unique" json:"password" binding:"required,min=6"`
	Email    string `gorm:"not null;unique" json:"email"  binding:"required,email"`
	Role     string `json:"role"`
}
