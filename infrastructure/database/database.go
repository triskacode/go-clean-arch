package database

import (
	"fmt"

	"github.com/triskacode/go-clean-arch/config"
	"gorm.io/gorm"
)

type DatabaseService struct {
	DB *gorm.DB
}

func New(cfg *config.Config) (dbs *DatabaseService) {
	db, err := newSqliteConnection(cfg.Database.Name)
	if err != nil {
		panic(fmt.Errorf("cannot connect database: %w", err))
	}

	defer func() {
		dbInstance, _ := db.DB()
		if err := dbInstance.Close(); err != nil {
			panic(fmt.Errorf("cannot close database connection: %w", err))
		}
	}()

	dbs = &DatabaseService{
		DB: db,
	}

	return
}
