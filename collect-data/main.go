package main

import (
	"log"

	"github.com/eufelipemateus/go-github-stats/collect-data/collects"
	config "github.com/eufelipemateus/go-github-stats/collect-data/configs"
	"github.com/eufelipemateus/go-github-stats/collect-data/database"
	"github.com/robfig/cron"
)

func main() {
	err := config.Load()
	if err != nil {
		log.Panicf("Erro on load config.toml")
	}
	database.OpenConnection()
	database.GenerateDB()
	/***/

	println("Startando crawler github...")
	c := cron.New()

	c.AddFunc("@every 15m",   func() { 
		collects.Collect(50, 100)
	})

	c.AddFunc("@every 1h",   func() { 
		collects.Collect(0, 50)
	})

	/*
	c.AddFunc("@every 1h30m", func() { 
		collects.Collect(75, 90)
	})

	c.AddFunc("@every 2h30m", func() { 
		collects.Collect(50, 75)
	})

	c.AddFunc("@every 4h",  func() { 
		collects.Collect(50, 75)
	})

	c.AddFunc("@every 6h", func() { 
		collects.Collect(0, 50)
	})

	*/

	c.Run()
}
