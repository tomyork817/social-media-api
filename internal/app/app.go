package app

import (
	"log"
	"log/slog"
	"os"
	"os/signal"
	"social-media-api/internal/controller/graph"
	"social-media-api/internal/infrastructure/inmemory"
	postgres2 "social-media-api/internal/infrastructure/postgres"
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

	var postRepo usecase.PostRepo
	var commentRepo usecase.CommentRepo

	switch cfg.Repository.Type {
	case "postgres":
		postRepo = postgres2.NewPostPostgres(pg)
		commentRepo = postgres2.NewCommentPostgres(pg)
	case "inmemory":
		postRepo = inmemory.NewPostInMemory()
		commentRepo = inmemory.NewCommentInMemory()
	default:
		log.Fatalf("no type of repo")
	}

	postUC := usecase.NewPostUseCase(postRepo)
	commentUC := usecase.NewCommentUseCase(commentRepo, postRepo)

	router := graph.NewRouter(postUC, commentUC)
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
}
