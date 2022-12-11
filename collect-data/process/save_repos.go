package process

import (
	"fmt"
	"log"

	"github.com/TwiN/go-color"
	"github.com/eufelipemateus/go-github-stats/collect-data/database/query"
	"github.com/eufelipemateus/go-github-stats/collect-data/models"
	"github.com/google/go-github/github"
)



func SaveRepos(repos []*github.Repository, CollectID uint, txCtx *query.IRepositoryDo) {

	log.Println(color.Ize(color.Green,"Savlando novos repositorios"))

	for i := range repos {
		name := repos[i].Name
		stars :=repos[i].StargazersCount
		fork :=repos[i].Fork
		forkCount :=repos[i].ForksCount
		if !*fork {
			reposData := models.Repository{Name: *name, CountStars: *stars, CountForks: *forkCount, CollectID: CollectID}
			(*txCtx).Create(&reposData) 
			log.Println(color.Ize(color.Yellow,fmt.Sprintf("Salvando respositorio \"%s\"...", *name )))
		}
	}
	log.Println(color.Ize(color.Green,"Todos novos repositorios salvos com sucesso."))

}