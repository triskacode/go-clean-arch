package database

import "gorm.io/gorm"

type DatabaseAdapter interface {
	GetConnection() *gorm.DB
	CloseConnection()
	Migrate(domain ...interface{})
}
