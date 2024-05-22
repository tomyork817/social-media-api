package main

import (
	"log"
	"social-media-api/internal/app"
	"social-media-api/pkg/config"
)

/*
migrate -database "postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable" -path db/migrations up
*/

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("unable to read config: %s", err)
	}

	app.Run(cfg)
}
