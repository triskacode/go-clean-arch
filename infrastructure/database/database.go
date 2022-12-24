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
	db, err := newSqliteConnection(cfg.Database.Sqlite.Name)
	if err != nil {
		panic(fmt.Errorf("cannot connect database: %w", err))
	}

	dbs = &DatabaseService{DB: db}
	return
}

func (dbs *DatabaseService) Migrate(domain ...interface{}) {
	if err := dbs.DB.AutoMigrate(domain...); err != nil {
		panic(fmt.Errorf("failed migrate database: %w", err))
	}
}

func (dbs *DatabaseService) CloseConnection() {
	conn, _ := dbs.DB.DB()
	if err := conn.Close(); err != nil {
		panic(fmt.Errorf("cannot close database connection: %w", err))
	}
}
