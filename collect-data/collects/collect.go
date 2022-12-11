package collects

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/TwiN/go-color"
	"github.com/eufelipemateus/go-github-stats/collect-data/database"
	"github.com/eufelipemateus/go-github-stats/collect-data/database/query"
	"github.com/eufelipemateus/go-github-stats/collect-data/models"
	"github.com/eufelipemateus/go-github-stats/collect-data/process"

	config "github.com/eufelipemateus/go-github-stats/collect-data/configs"
)

func Collect(relevance int, relevanceEnd int) {
	limit := config.GetCollectLimit()

	ctx := context.Background()
	userQuery := query.User
	users, _ := userQuery.WithContext(ctx).Where(
		userQuery.Relevance.Between(relevance, relevanceEnd),
	).Order(
		userQuery.UpdatedAt,
	).Limit(limit).Find()

	for e := range users{
		go CollectUser(users[e])
	}
}

func CollectUser(user *models.User){

	username := user.Username

	ctx := context.Background()
	q := query.Use(database.DB)
	tx := q.Begin()
	txCtx := tx.WithContext(ctx)


	collect := query.Collect
	collect_new := models.Collect{ UserID:  user.ID, CollectStartedAt: time.Now()}

	txCtx.Collect.Create(&collect_new)
	
	repos, countStars := GetUserRepos(username)
	collect_new.CountRepos = len(repos)
	collect_new.CountStars = countStars
	go process.SaveRepos(repos, collect_new.ID, &txCtx.Repository)
	tx.SavePoint("repos")

	langs, _ := GetUserRepositoryLanguages(username, repos)
	collect_new.CountLanguages = len(langs)
	go process.SaveLangs(langs, collect_new.ID, &txCtx.Lang )
	tx.SavePoint("langs")

	followers := GetFollowers(username)
	collect_new.CountFollowers = len(followers)
	newUsers := process.SaveUser(followers, &txCtx.User)
	process.SaveFollower( collect_new, newUsers, &txCtx.Collect)
	tx.SavePoint("followers")

	collect_new.FinishedAt = time.Now()

	txCtx.Collect.Select(
		collect.CountLanguages,
		collect.CountRepos,
		collect.CountFollowers,
		collect.FinishedAt,
		collect.CountStars,
	).Updates(&collect_new)
	tx.SavePoint("collect")

	process.CalculateRelavance(
		username,
		collect_new,
		&txCtx.User,
		&txCtx.Collect,
	)

	tx.Commit()
	log.Println(color.Ize(color.Blue,fmt.Sprintf("Todos dados do usuario \"%s\" foram coletados e processados.",username )))
}