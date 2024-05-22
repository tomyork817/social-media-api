package app

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"log"
	"net/http"
	"os"
	"social-media-api/config"
	"social-media-api/internal/controller/graph"
	"social-media-api/internal/controller/graph/generated"
	"social-media-api/internal/infrastructure/inmemory"
	"social-media-api/internal/usecase"
	"social-media-api/pkg/postgres"
)

func Run(cfg *config.Config) {
	repo := inmemory.NewPostInMemory()
	uc := usecase.NewPostUseCase(repo)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{PostUseCase: uc}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", cfg.HTTP.Port)

	pg, err := postgres.New(cfg.Postgres)
	if err != nil {
		log.Fatalf("can't connect to db: %s", err)
	}
	defer pg.Close()

	var greeting string
	err = pg.Pool.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)
}
