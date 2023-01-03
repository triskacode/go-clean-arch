package database

import (
	"fmt"

	"github.com/triskacode/go-clean-arch/config"
	"gorm.io/gorm"
)

type databaseService struct {
	db     *gorm.DB
	config *config.Config
}

func NewDatabaseService(cfg *config.Config) (dbSvc *databaseService) {
	db, err := newSqliteConnection(cfg.Database.Sqlite.Name)
	if err != nil {
		panic(fmt.Errorf("cannot connect database: %w", err))
	}

	dbSvc = new(databaseService)
	dbSvc.db = db
	dbSvc.config = cfg

	return
}

func (dbSvc databaseService) GetConnection() *gorm.DB {
	return dbSvc.db
}

func (dbSvc databaseService) CloseConnection() {
	conn, _ := dbSvc.GetConnection().DB()
	if err := conn.Close(); err != nil {
		panic(fmt.Errorf("cannot close database connection: %w", err))
	}
}

func (dbSvc databaseService) Migrate(domain ...interface{}) {
	if err := dbSvc.GetConnection().AutoMigrate(domain...); err != nil {
		panic(fmt.Errorf("failed migrate database: %w", err))
	}
}
