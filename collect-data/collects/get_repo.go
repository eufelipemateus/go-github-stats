package collects

import (
	"context"
	"fmt"
	"log"

	"github.com/TwiN/go-color"
	"github.com/google/go-github/github"
)


func GetUserRepos (username string) ([]*github.Repository,  int) {
	

	var reposUser []*github.Repository
	var page = 1;
	var starsRepos = 0 
	log.Println(color.Ize(color.Green,"Iniciando a coleta de repos do usuario: "+username))
	ctx := context.Background()

	auth := GithubAuth(ctx)

	for{
		client :=github.NewClient(auth[page%len(auth)])
		RateLimit(client, ctx)
		opt := &github.RepositoryListOptions{Type: "owner", ListOptions: github.ListOptions{PerPage: 100, Page: page } }
		repos, resp, err := client.Repositories.List(ctx, username, opt)
		if err != nil {
			log.Println(color.Ize(color.Red, err.Error()))
			continue
		}
		reposUser = append(reposUser, repos... )
		if resp.NextPage <= 0 {
			break
		}
		page = resp.NextPage
	}
	log.Println(color.Ize(color.Blue,fmt.Sprintf(" Coletado \"%d\" repositorio.",len(reposUser))))
	log.Println(color.Ize(color.Green,fmt.Sprintf(" Coletado todos os Repositorios do Usuario \"%s\" com sucesso.",username)))
	
	
	log.Println(color.Ize(color.Blue,fmt.Sprintf(" Somando Stars do Usuario \"%s\" ...",username)))
	for i := range reposUser {
		starsRepos += *reposUser[i].StargazersCount
	}
	
	log.Println(color.Ize(color.Green,fmt.Sprintf(" Todas Stars do Usuario \"%s\" foram somadas com sucesso.",username)))
	return reposUser, starsRepos

}