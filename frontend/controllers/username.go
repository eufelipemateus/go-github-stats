package controllers

import (
	"context"
	"errors"
	"fmt"
	"sort"

	"github.com/eufelipemateus/go-github-stats/frontend/database/query"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetUser(c *fiber.Ctx) error {

	ctx := context.Background()
	user := query.User
	collect := query.Collect
	lang := query.Lang
	repo := query.Repository

	username := c.Params("user")

	userData, err := user.WithContext(ctx).Where(user.Username.Eq(username), user.Show.Is(true)).First()

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(404).Render("errors/404", "")
	}

	collectData, _ := collect.WithContext(ctx).Preload(collect.Followers).Where(collect.UserID.Eq(userData.ID)).First()
	langData, _ := lang.WithContext(ctx).Where(lang.CollectID.Eq(collectData.ID)).Order(lang.Percent.Desc()).Find()
	reposData, _ := repo.WithContext(ctx).Where(
		repo.CollectID.Eq(collectData.ID),
	).Order(
		repo.CountStars.Desc(),
		repo.CountForks.Desc(),
	).Limit(4).Find()


	followers := collectData.Followers
	total := len(followers)

	if  total > 12 {
		total = 12
	}

	sort.SliceStable(followers, func(i, j int) bool {
		return followers[i].Relevance > followers[j].Relevance &&   followers[i].Show
	})

	followers = followers[0:total]

	return c.Render("users", fiber.Map{
		"Title":     fmt.Sprintf("Usuario, %s", userData.Username),
		"Username":  userData.Username,
		"Relevance": userData.Relevance,
		"LastCollect": fmt.Sprintf("%02d/%02d/%d %02d:%02d:%02d BRT\n",
			collectData.FinishedAt.Day(), collectData.FinishedAt.Month(), collectData.FinishedAt.Year(),
			collectData.FinishedAt.Hour(), collectData.FinishedAt.Minute(), collectData.FinishedAt.Second()),
		"TotalStarts":  collectData.CountStars,
		"TotalRepos":   collectData.CountRepos,
		"TotalLangs":   collectData.CountLanguages,
		"Langs":        langData,
		"Repositories": reposData,
		"Followers":   followers,
	})
}
