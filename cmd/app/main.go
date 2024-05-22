package main

import (
	"log"
	"social-media-api/config"
	"social-media-api/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("unable to read config: %s", err)
	}

	app.Run(cfg)
}
