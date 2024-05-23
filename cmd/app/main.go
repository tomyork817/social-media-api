package main

import (
	"log"
	"log/slog"
	"os"
	"social-media-api/internal/app"
	"social-media-api/pkg/config"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("unable to read config: %s", err)
	}

	app.Run(cfg)
}
