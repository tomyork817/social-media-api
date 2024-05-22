package app

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"log"
	"net/http"
	"social-media-api/config"
	"social-media-api/internal/controller/graph"
	"social-media-api/internal/controller/graph/generated"
	"social-media-api/internal/infrastructure/inmemory"
	"social-media-api/internal/usecase"
)

func Run(cfg *config.Config) {
	repo := inmemory.NewPostInMemory()
	uc := usecase.NewPostUseCase(repo)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{PostUseCase: uc}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", cfg.HTTP.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.HTTP.Port, nil))
}
