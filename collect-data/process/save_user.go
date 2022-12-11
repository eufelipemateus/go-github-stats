package process

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/TwiN/go-color"
	"github.com/eufelipemateus/go-github-stats/collect-data/database/query"
	"github.com/eufelipemateus/go-github-stats/collect-data/models"
	"github.com/google/go-github/github"
	"gorm.io/gorm"
)



func SaveUser(users []*github.User, txCtx  *query.IUserDo) []*models.User{
	ctx := context.Background()
	userQuery := query.User

	var followers []* models.User

	log.Println(color.Ize(color.Green,"Savlando novos usuarios"))

	for i := range users {
		username := *users[i].Login
		userData := models.User{Username: username, Relevance: 100}
		userF, err := userQuery.WithContext(ctx).Where(userQuery.Username.Eq(username)).First()

		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println(color.Ize(color.Yellow,fmt.Sprintf("%d. Savlando usuario \"%s\"...", i+1, username )))
			(*txCtx).Create(&userData) 
			followers = append(followers, &userData)
			continue
		}
		followers = append(followers, userF)
		log.Println(color.Ize(color.Blue,fmt.Sprintf("%d.  Usuario  \"%s\" Já existe.", i+1, username )))		
	}

	log.Println(color.Ize(color.Green,"Todos novos usuarios salvos com sucesso."))
	return followers
}