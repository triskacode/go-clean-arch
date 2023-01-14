package main

import (
	"github.com/triskacode/go-clean-arch/config"
	"github.com/triskacode/go-clean-arch/domain/entity"
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

	dbSvc.Migrate(&entity.Article{}, &entity.Author{})

	authorMod := author.NewModule(author.ModuleDeps{
		App: appSvc.GetApp(),
		DB:  dbSvc.GetConnection(),
	})

	articleMod := article.NewModule(article.ModuleDeps{
		App:              appSvc.GetApp(),
		DB:               dbSvc.GetConnection(),
		AuthorRepository: authorMod.GetAuthorRepository(),
	})

	authorMod.InitializeRoute()
	articleMod.InitializeRoute()

	appSvc.Run()
}
