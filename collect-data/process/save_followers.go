package process

import (
	"log"

	"github.com/TwiN/go-color"
	"github.com/eufelipemateus/go-github-stats/collect-data/database/query"
	"github.com/eufelipemateus/go-github-stats/collect-data/models"
)


func SaveFollower( collect models.Collect, users []*models.User , txCtx  *query.ICollectDo){
	log.Println(color.Ize(color.Green,"Salvando Followers..."))
		collect.Followers = append(collect.Followers, users... )

		(*txCtx).Updates(&collect )
	
	log.Println(color.Ize(color.Green,"Todos os Followers foram salvos com sucesso."))
}