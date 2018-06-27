package models

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // SQLite
)

// Leave - Leave Model
type Leave struct {
	gorm.Model

	UserRefer uint64
	Reason    string
	LeaveDate time.Time
}

// LeaveResponse - Leave Response Model
type LeaveResponse struct {
	LeaveID   string    `json:"id"`
	Reason    string    `json:"reason"`
	LeaveDate time.Time `json:"date"`
}
