package graph

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"net/http"
	"social-media-api/internal/controller/graph/generated"
	"social-media-api/internal/usecase"
)

type Router struct {
	GraphQLServer *handler.Server
	Multiplexer   *http.ServeMux
}

func NewRouter(post usecase.Post, comment usecase.Comment) *Router {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{PostUseCase: post, CommentUseCase: comment},
	}))
	mux := http.NewServeMux()

	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mux.HandleFunc("POST /query", srv.ServeHTTP)

	return &Router{
		GraphQLServer: srv,
		Multiplexer:   mux,
	}
}
