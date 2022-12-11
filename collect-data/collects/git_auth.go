package collects

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/TwiN/go-color"
	config "github.com/eufelipemateus/go-github-stats/collect-data/configs"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)


func RateLimit(client  *github.Client, ctx context.Context){

	_, resp, _ := client.RateLimits(ctx)
	if  resp.Rate.Remaining <= 0 {
		log.Println(color.Ize(color.Red,"Rate Limit Atigido"))
		endDate := time.Until(resp.Rate.Reset.Time) 
		log.Println(color.Ize(color.Blue, "Aguardando liberação do Rate Limit..."))
		time.Sleep(time.Duration(time.Duration(endDate).Milliseconds()))
	}
}

func GithubAuth(ctx context.Context)  []*http.Client{
	tokens := config.GetGithubTokens()

	var authHttp [] *http.Client

	for i := range tokens {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: tokens[i]  },
		)
		auth := oauth2.NewClient(ctx, ts)
		authHttp = append(authHttp, auth )
	}
	return authHttp
}