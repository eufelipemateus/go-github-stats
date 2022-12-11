package database

import (
	"github.com/eufelipemateus/go-github-stats/frontend/models"
	"gorm.io/gen"
)

func GenerateDB() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./database/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	g.UseDB(DB)

	g.ApplyBasic(models.User{}, models.Lang{}, models.Collect{}, models.Repository{})

	g.Execute()
	DB.AutoMigrate(models.User{}, models.Collect{}, models.Lang{}, models.Repository{})
}
