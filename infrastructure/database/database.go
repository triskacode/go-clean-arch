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
		panic(fmt.Errorf("fatal error connect database: %w", err))
	}

	dbs = &DatabaseService{
		DB: db,
	}

	return
}
