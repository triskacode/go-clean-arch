package main

import (
	"github.com/triskacode/go-clean-arch/config"
	"github.com/triskacode/go-clean-arch/domain"
	"github.com/triskacode/go-clean-arch/infrastructure/database"
	"github.com/triskacode/go-clean-arch/infrastructure/http"
	author "github.com/triskacode/go-clean-arch/modules/author/bootstrap"
)

func main() {
	cfg := config.New()
	appSvc := http.NewHttpService(cfg)
	dbSvc := database.NewDatabaseService(cfg)
	defer dbSvc.CloseConnection()

	dbSvc.Migrate(&domain.Article{}, &domain.Author{})

	authorMod := author.NewModule(appSvc.GetApp(), dbSvc.GetConnection())
	authorMod.InitializeRoute()

	appSvc.Run()
}
