package collects

import (
	"context"
	"fmt"
	"log"

	"github.com/TwiN/go-color"
	"github.com/google/go-github/github"
)


func GetFollowers (username string) []*github.User{

	var followersUser []*github.User
	var page = 1;
	log.Println(color.Ize(color.Green,"Iniciando a coleta de followers do usuario: "+username))
	ctx := context.Background()
	auth := GithubAuth(ctx)

	for{
		client :=github.NewClient(auth[page%len(auth)])
		RateLimit(client, ctx)

		opt := &github.ListOptions{PerPage: 100, Page: page } 
		users, resp, err := client.Users.ListFollowers(ctx, username, opt )

		if err != nil {
			log.Println(color.Ize(color.Red, err.Error()))
		}
		followersUser = append(followersUser, users... )
		if resp.NextPage <= 0 {
			break
		}
		page = resp.NextPage
	}
	log.Println(color.Ize(color.Blue,fmt.Sprintf(" Coletado \"%d\" followers.",len(followersUser))))
	log.Println(color.Ize(color.Green,fmt.Sprintf(" Finalizado a Coleta de todos os Followers do Usuario \"%s\" com sucesso.",username)))
	return  followersUser
}