package process

import (
	"math"

	"github.com/eufelipemateus/go-github-stats/collect-data/database/query"
	"github.com/eufelipemateus/go-github-stats/collect-data/models"
)



func CalculateRelavance(
	username string,
	collect models.Collect,
	txCtxUser *query.IUserDo,
	txtCtxCollect *query.ICollectDo,
){
	userQuery := query.User
/*collectQuery := query.Collect

	ctx := context.Background()

	var collectTotal struct {
		TotalLangs  int64
		TotalFollowers int64
		TotalRepos int64
		TotalStarts int64
	}*/

    countFollowres := int64(collect.CountFollowers)
    countLanguages := int64(collect.CountLanguages)
	countRepos := int64(collect.CountRepos)
	countStars := int64(collect.CountStars)


	/*collectQuery.WithContext(ctx).Select(
		collectQuery.CountLanguages.Sum().As("total_langs"),
		collectQuery.CountFollowers.Sum().As("total_followers"),
		collectQuery.CountRepos.Sum().As("total_repos"),
		collectQuery.CountStars.Sum().As("total_starts"),
	).Scan(&collectTotal)

    total := int64( collectTotal.TotalLangs + collectTotal.TotalRepos + collectTotal.TotalStarts + collectTotal.TotalFollowers
	average:= float64(0)
	relevance:= 100 

	if total > 0{
		average = float64(
			(countLanguages*5)+(countRepos*5)+(countStars*20)+(countFollowres*10)/
			(total*40))
			relevance = int(math.Log(average))
	}*/

	average := float64((countLanguages*5)+(countRepos*5)+(countStars*20)+(countFollowres*10)/40)
	relevance := int(math.Log(average))

	if relevance <  0 {
		relevance = 0
	} 

	userData := models.User{Show: true, Relevance: relevance}
	(*txCtxUser).Select(
		userQuery.Relevance,
		userQuery.Show,
	).Where(
		userQuery.Username.Eq(username),
	).Updates(&userData)

}