package models

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // SQLite
)

// Leave - Leave Model
type Leave struct {
	gorm.Model `json:"-"`

	UserRefer uint64    `json:"-"`
	Reason    string    `json:"reason"`
	LeaveDate time.Time `json:"date"`
}
