package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init(runMode string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("data.sqlite3."+runMode), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}

func MigrateDB() {
	DB.AutoMigrate(
		&Todo{},
	)
}
