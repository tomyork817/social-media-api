package app

import (
	"log"
	"log/slog"
	"os"
	"os/signal"
	"social-media-api/internal/controller/graph"
	"social-media-api/internal/infrastructure/inmemory"
	"social-media-api/internal/usecase"
	"social-media-api/pkg/config"
	"social-media-api/pkg/httpserver"
	"social-media-api/pkg/postgres"
	"syscall"
)

func Run(cfg *config.Config) {
	pg, err := postgres.New(cfg.Postgres)
	if err != nil {
		log.Fatalf("can't connect to db: %s", err)
	}
	defer pg.Close()

	repo := inmemory.NewPostInMemory()
	uc := usecase.NewPostUseCase(repo)
	router := graph.NewRouter(uc)

	httpServer := httpserver.New(router.Multiplexer, cfg.HTTP)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", cfg.HTTP.Port)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		slog.Info(s.String())
	case err = <-httpServer.Notify():
		slog.Error(err.Error())
	}

	slog.Info("shutting down")

	if err = httpServer.Shutdown(); err != nil {
		slog.Error(err.Error())
	}

	/*	var greeting string
		err = pg.Pool.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
		if err != nil {
			fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(greeting)*/
}
