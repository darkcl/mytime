package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // SQLite
)

// User - User Model
type User struct {
	gorm.Model

	Username string `json:"userName"`
	Password []byte `gorm:"type:varchar(255); not null"`

	Leave Leave `gorm:"foreignkey:UserRefer"`
}

// JwtToken - Token Model
type JwtToken struct {
	Token string `json:"token"`
}
