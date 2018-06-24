package db

import (
	"fmt"

	"github.com/darkcl/mytime/models"
	"github.com/jinzhu/gorm"

	// Using SQLite3 as our development database
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var conn *gorm.DB
var dbError error

// Init - Initalize Database
func Init() {
	conn, dbError = gorm.Open("sqlite3", "time.sqlite")
	if dbError != nil {
		panic("Failed to connect database")
	}
	conn.AutoMigrate(&models.User{})
	conn.AutoMigrate(&models.Leave{})
	fmt.Println("Connected To Database")
}

// GetDB - Get Database Connection
func GetDB() *gorm.DB {
	return conn
}

// CloseDB - Close Database Connection
func CloseDB() {
	conn.Close()
}
