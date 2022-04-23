package storage

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDb() (*gorm.DB, error) {
	return newSQLite()
}

func newSQLite() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db, nil
}
