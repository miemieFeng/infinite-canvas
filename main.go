package main

import (
	"log"

	"github.com/miemieFeng/infinite-canvas/config"
	"github.com/miemieFeng/infinite-canvas/router"
	"github.com/miemieFeng/infinite-canvas/service"
)

func main() {
	if err := config.Load(); err != nil {
		log.Fatal(err)
	}
	if err := service.EnsureDefaultAdmin(); err != nil {
		log.Fatal(err)
	}
	service.StartPromptSyncScheduler()
	log.Fatal(router.New().Run(":" + config.Cfg.Port))
}
