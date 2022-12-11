package database

import (
	"fmt"

	config "github.com/eufelipemateus/go-github-stats/frontend/configs"
	"github.com/eufelipemateus/go-github-stats/frontend/database/query"

	//"github.com/eufelipemateus/calculadora-trabalhista-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func OpenConnection() {
	conf := config.GetDB()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable ",
		conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		fmt.Printf("Erro while try connect to database.")
	}

	DB = database

	query.SetDefault(DB)
}
