package main

import (
	"github.com/triskacode/go-clean-arch/config"
	"github.com/triskacode/go-clean-arch/domain"
	"github.com/triskacode/go-clean-arch/infrastructure/database"
	"github.com/triskacode/go-clean-arch/infrastructure/http"
	article "github.com/triskacode/go-clean-arch/modules/article/bootstrap"
	author "github.com/triskacode/go-clean-arch/modules/author/bootstrap"
)

func main() {
	cfg := config.New()
	appSvc := http.NewHttpService(cfg)
	dbSvc := database.NewDatabaseService(cfg)
	defer dbSvc.CloseConnection()

	dbSvc.Migrate(&domain.Article{}, &domain.Author{})

	authorDeps := author.ModuleDeps{
		App: appSvc.GetApp(),
		DB:  dbSvc.GetConnection(),
	}
	authorMod := author.NewModule(authorDeps)

	articleDeps := article.ModuleDeps{App: appSvc.GetApp()}
	articleMod := article.NewModule(articleDeps)

	authorMod.InitializeRoute()
	articleMod.InitializeRoute()

	appSvc.Run()
}
