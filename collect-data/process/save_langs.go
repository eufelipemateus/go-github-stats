package process

import (
	"log"

	"github.com/TwiN/go-color"
	"github.com/eufelipemateus/go-github-stats/collect-data/database/query"
	"github.com/eufelipemateus/go-github-stats/collect-data/models"
)


func SaveLangs(langs []models.ApiLang, id uint,  txCtx *query.ILangDo){
	log.Println(color.Ize(color.Green,"Salvando langs..."))
	for i := range langs {
		lang := &models.Lang{
			Language: langs[i].Language,
			Weight: langs[i].Weight,
			Percent: float32(langs[i].Percent),
			CollectID: id,
		}
		(*txCtx).Create(lang )
	}
	log.Println(color.Ize(color.Green,"Langs salvas com sucesso."))
}