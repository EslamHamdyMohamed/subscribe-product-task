package main

import (
	"github.com/go-co-op/gocron"
	"github.com/subosito/gotenv"
	"subscribe-product/app/cron"
	"subscribe-product/database"
	"subscribe-product/providers"
	"time"
)

func main() {
	err := gotenv.Load()
	if err != nil {
		panic(err)
	}
	database.ConnectToDatabase()
	database.MigrateAll()

	s := gocron.NewScheduler(time.UTC)
	cron.Jobs(s)
	s.StartAsync()

	providers.Run()
}
