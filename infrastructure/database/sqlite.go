package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func newSqliteConnection(dsn string) (db *gorm.DB, err error) {
	db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})

	return
}
