package collects

import (
	"context"
	"fmt"
	"log"

	"github.com/TwiN/go-color"
	"github.com/eufelipemateus/go-github-stats/collect-data/models"
	"github.com/google/go-github/github"
)




func GetUserRepositoryLanguages (username string, repos []*github.Repository) ([] models.ApiLang, int  )  {
	var userLangsMap map[string]int = make(map[string]int)
	var userLangs []models.ApiLang
	var totaLangsWeight int

	log.Println(color.Ize(color.Green,"Iniciando a coleta de langs do usuario: "+username))
	ctx := context.Background()

	auth := GithubAuth(ctx)

	for i := range repos {
		client :=github.NewClient(auth[i%len(auth)])
		repository := *repos[i].Name
		RateLimit(client, ctx)
		langs, _, err := client.Repositories.ListLanguages(ctx, username, repository)
		if err != nil {
			log.Println(color.Ize(color.Red, err.Error()))
			continue
		}
		for k, s  := range langs {
			userLangsMap[k] = userLangsMap[k] + s
		}	

		//log.Println(color.Ize(color.Yellow,fmt.Sprintf("%d. Todas as linguagens do repositorio \"%s\" foram coletadas.", i+1, repository)))
	}
	log.Println(color.Ize(color.Blue,fmt.Sprintf(" Coletado \"%d\" langs.",len(userLangs))))
	log.Println(color.Ize(color.Blue,fmt.Sprintf("Somando todos as linguagens usuario \"%s\"...", username)))

	for i, s := range userLangsMap {
		lang := models.ApiLang{Language: i, Weight: s}
		userLangs = append(userLangs, lang)
	}
	for  i := range userLangs{
		totaLangsWeight += userLangs[i].Weight
	}
	for  i := range userLangs{
		percent := float64(userLangs[i].Weight) / float64(totaLangsWeight) * 100
		userLangs[i].Percent = percent
	}

	log.Println(color.Ize(color.Green,fmt.Sprintf("Todas as linguagens do usuario \"%s\" foram processadas.", username)))
	return userLangs, totaLangsWeight
}
