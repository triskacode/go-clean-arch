package database

import (
	"fmt"

	"github.com/triskacode/go-clean-arch/config"
	"gorm.io/gorm"
)

type databaseService struct {
	db *gorm.DB
}

func NewConnection(cfg *config.Config) (dbs *databaseService) {
	db, err := newSqliteConnection(cfg.Database.Sqlite.Name)
	if err != nil {
		panic(fmt.Errorf("cannot connect database: %w", err))
	}

	dbs = &databaseService{db: db}
	return
}

func (dbs databaseService) GetConnection() *gorm.DB {
	return dbs.db
}

func (dbs databaseService) CloseConnection() {
	conn, _ := dbs.GetConnection().DB()
	if err := conn.Close(); err != nil {
		panic(fmt.Errorf("cannot close database connection: %w", err))
	}
}

func (dbs databaseService) Migrate(domain ...interface{}) {
	if err := dbs.GetConnection().AutoMigrate(domain...); err != nil {
		panic(fmt.Errorf("failed migrate database: %w", err))
	}
}
